package core

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
)

const sessionMetricsMarker = "/* regravity:session-metrics */"
const legacySessionMetricsMarker = "/* antigravity-i18n:session-metrics */"
const metricsLocalePlaceholder = "METRICS_LOCALE_PLACEHOLDER"

func (s *LocaleStore) addSessionMetrics(patch *patchFragments, loc *Locale) error {
	main, err := fs.ReadFile(s.patchesFS, s.patchDir+"/metrics-main.jsfrag")
	if err != nil {
		return fmt.Errorf("cannot read the session metrics reader: %w", err)
	}
	preload, err := fs.ReadFile(s.patchesFS, s.patchDir+"/metrics-preload.jsfrag")
	if err != nil {
		return fmt.Errorf("cannot read the session metrics view: %w", err)
	}
	dictionary, err := fs.ReadFile(s.patchesFS, s.patchDir+"/metrics-labels.json")
	if err != nil {
		return fmt.Errorf("cannot read session metrics labels: %w", err)
	}
	var labels map[string]map[string]string
	if err := json.Unmarshal(dictionary, &labels); err != nil {
		return fmt.Errorf("cannot parse session metrics labels: %w", err)
	}
	selected := labels[loc.Language]
	if selected == nil {
		return fmt.Errorf("no session metrics labels for %s", loc.Language)
	}
	selected["language"] = loc.HTMLLang
	if selected["language"] == "" {
		selected["language"] = loc.Language
	}
	data, err := json.Marshal(selected)
	if err != nil {
		return err
	}
	view, n := replaceFirst(string(preload), metricsLocalePlaceholder, string(data))
	if n != 1 {
		return fmt.Errorf("the session metrics view has no %s", metricsLocalePlaceholder)
	}
	// Keep the modules inside the existing patch block so changing languages
	// or switching the feature off replaces the entire previous block.
	end := blockSentinels[0][1]
	patch.main = strings.Replace(patch.main, end, string(main)+"\n"+end, 1)
	patch.preload = strings.Replace(patch.preload, end, sessionMetricsMarker+"\n"+view+"\n"+end, 1)
	return nil
}
