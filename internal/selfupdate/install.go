package selfupdate

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

// Each installation has one receipt. The random ID identifies exactly the two
// sibling files allocated for this attempt, including after a process restart.
type receipt struct {
	ID   string
	From string
	To   string
}

func receiptPath(exe string) string        { return exe + ".regravity-update.json" }
func (r receipt) staged(exe string) string { return exe + "." + r.ID + ".new.exe" }
func (r receipt) backup(exe string) string { return exe + "." + r.ID + ".old.exe" }

func newReceipt(exe, current, target string) (receipt, error) {
	var token [16]byte
	rand.Read(token[:])
	r := receipt{ID: hex.EncodeToString(token[:]), From: versionTag(current), To: target}
	// Publish the complete record at once so an interrupted write cannot
	// leave an unreadable receipt that blocks every subsequent update.
	f, err := os.CreateTemp(filepath.Dir(exe), ".regravity-receipt-*.tmp")
	if err != nil {
		return r, fmt.Errorf("无法在工具所在目录创建更新记录：%w", err)
	}
	defer os.Remove(f.Name())
	err = json.NewEncoder(f).Encode(r)
	if err == nil {
		err = f.Sync()
	}
	if err = errors.Join(err, f.Close()); err != nil {
		return r, err
	}
	return r, moveNewName(f.Name(), receiptPath(exe))
}

func removeIfPresent(path string) error {
	err := os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

// Called under the per-executable lock. An interrupted download is disposable.
// A backup can be removed once the target version (or a subsequent manual
// upgrade) has reached the GUI.
func cleanup(exe, current string) error {
	data, err := os.ReadFile(receiptPath(exe))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	var r receipt
	if err = json.Unmarshal(data, &r); err != nil {
		return err
	}
	id, err := hex.DecodeString(r.ID)
	if err != nil || len(id) != 16 || !semver.IsValid(r.From) || !semver.IsValid(r.To) || semver.Compare(r.From, r.To) >= 0 {
		return fmt.Errorf("更新记录无效，请检查 %s", receiptPath(exe))
	}
	tag := versionTag(current)
	switch {
	case semver.IsValid(tag) && semver.Compare(tag, r.To) >= 0:
		if err = removeIfPresent(r.backup(exe)); err != nil {
			return err
		}
	case tag == r.From:
		if _, err := os.Lstat(r.backup(exe)); !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("上次更新保留的旧版位于 %s", r.backup(exe))
		}
	default:
		return fmt.Errorf("上次更新（%s → %s）的记录位于 %s", r.From, r.To, receiptPath(exe))
	}
	if err = removeIfPresent(r.staged(exe)); err != nil {
		return err
	}
	return removeIfPresent(receiptPath(exe))
}

// WindowReady runs in the background only after the frontend displays its version.
// The old process may still have the backup mapped while its window is closing.
func WindowReady(current string) error {
	if err := signalReady(); err != nil {
		return err
	}
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	if _, err = os.Stat(receiptPath(exe)); errors.Is(err, os.ErrNotExist) {
		return nil
	}
	for attempt := 0; attempt < 26; attempt++ {
		var unlock func()
		unlock, err = lockExecutable(exe)
		if err == nil {
			err = cleanup(exe, current)
			unlock()
		}
		if err == nil || attempt == 25 {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	return err
}

func verifyVersion(ctx context.Context, exe, version string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, exe, "--version")
	hideProbeWindow(cmd)
	output, err := cmd.Output()
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if err != nil {
		return err
	}
	expected := "regravity " + strings.TrimPrefix(version, "v")
	if strings.TrimSpace(string(output)) != expected {
		return fmt.Errorf("程序版本不符，预期 %s。", expected)
	}
	return nil
}

const readyEnv = "REGRAVITY_UPDATE_READY_EVENT"

const RestartDirEnv = "REGRAVITY_RESTART_APP_DIR"

func (r *Release) Install(ctx context.Context, current, appDir string, report func(string, string), beginSwap func() error) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	unlock, err := lockExecutable(exe)
	if err != nil {
		return err
	}
	defer unlock()
	// A second GUI opened before an update may still report the older version.
	if err := verifyVersion(ctx, exe, current); err != nil {
		return fmt.Errorf("请重新打开当前工具再更新：%w", err)
	}
	if err := cleanup(exe, current); err != nil {
		return err
	}
	tx, err := newReceipt(exe, current, r.Version)
	if err != nil {
		return err
	}
	err = r.download(ctx, tx.staged(exe), report)
	if err == nil {
		report("正在校验版本", r.Version)
		err = verifyVersion(ctx, tx.staged(exe), r.Version)
	}
	if err == nil {
		err = beginSwap()
	}
	if err == nil {
		report("正在重启工具", r.Version)
		err = replaceAndStart(exe, tx, func() error {
			cmd := exec.Command(exe)
			cmd.Env = append(os.Environ(), RestartDirEnv+"="+appDir)
			return startAndWaitReady(cmd, tx.ID, 30*time.Second)
		})
	}
	if err != nil {
		// Keep the receipt if rollback failed. That preserves the runnable backup.
		return errors.Join(err, cleanup(exe, current))
	}
	return nil
}

func replaceAndStart(exe string, tx receipt, start func() error) error {
	old, staged := tx.backup(exe), tx.staged(exe)
	if err := moveNewName(exe, old); err != nil {
		return err
	}
	err := moveNewName(staged, exe)
	if err == nil {
		if err = start(); err == nil {
			return nil
		}
		// Clear the original name before restoring the backup.
		if moveErr := moveNewName(exe, staged); moveErr != nil {
			return fmt.Errorf("新版启动失败（%v），可直接打开旧版 %s：%w", err, old, moveErr)
		}
	}
	if restoreErr := moveNewName(old, exe); restoreErr != nil {
		return fmt.Errorf("更新失败（%v），可直接打开旧版 %s：%w", err, old, restoreErr)
	}
	return fmt.Errorf("已恢复旧版：%w", err)
}
