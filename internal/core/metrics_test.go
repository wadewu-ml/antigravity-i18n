package core

import (
	"context"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestSessionMetricsInBrowser(t *testing.T) {
	store := newTestService(t).store
	loc, err := store.LoadLocale(DefaultLocale)
	if err != nil {
		t.Fatal(err)
	}
	preload, err := store.buildPreloadFragment(loc)
	if err != nil {
		t.Fatal(err)
	}
	patch := patchFragments{preload: preload}
	if err := store.addSessionMetrics(&patch, loc); err != nil {
		t.Fatal(err)
	}
	html := readTestFile(t, "testdata/metrics.html")
	runBrowserChecks(t, strings.Replace(html, "<!-- ENGINE -->", "<script>"+patch.preload+"</script>", 1))
}

func TestSessionMetricsSQLite(t *testing.T) {
	node, err := exec.LookPath("node")
	if err != nil {
		t.Skip("Node is needed only for the SQLite integration test")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := exec.CommandContext(ctx, node, "-e", "require('node:sqlite')").Run(); err != nil {
		t.Skip("the SQLite integration test needs a Node version with node:sqlite")
	}
	cmd := exec.CommandContext(ctx, node, "testdata/metrics-node.js", t.TempDir())
	hideWindow(cmd)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("session metrics SQLite test: %v\n%s", err, output)
	}
}
