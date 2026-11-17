package core

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"
)

// ASAR permits empty files to share the next body's offset.
func TestRepackKeepsSharedEmptyOffsets(t *testing.T) {
	header := pickleWrapJSON([]byte(`{"files":{"body":{"size":1,"offset":"0"},"empty":{"size":0,"offset":"0"}}}`))
	var lead [8]byte
	binary.LittleEndian.PutUint32(lead[:4], 4)
	binary.LittleEndian.PutUint32(lead[4:], uint32(len(header)))
	dir := t.TempDir()
	src := filepath.Join(dir, "original.asar")
	if err := os.WriteFile(src, append(append(lead[:], header...), 'x'), 0o644); err != nil {
		t.Fatal(err)
	}
	dst := filepath.Join(dir, "repacked.asar")
	if err := WritePatched(src, dst, nil); err != nil {
		t.Fatal(err)
	}
	if fileHash(t, dst) != fileHash(t, src) {
		t.Fatal("repacking moved the empty file away from its original offset")
	}
}
