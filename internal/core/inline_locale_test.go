package core

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func inlineLocaleFixture() *Locale {
	index := 2
	return &Locale{
		Language: "ja", Name: "日本語", Text: map[string]string{"Open": "開く"},
		InlineMessages: []InlineMessage{{
			ID: "docs", Selector: "p:has(>a)",
			Source: "Open {1} for {2}.", Template: "{1} を {2} のために開く。",
			Slots: []InlineSlot{{Selector: "a", Text: "Docs"}, {TextIndex: &index}},
		}},
	}
}

// Slot order preserves controls and arbitrary names in their original DOM
// positions. Reject malformed packs before they can rewrite an inline message.
func TestInlineMessageValidation(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*Locale)
		want string
	}{
		{"empty-id", func(l *Locale) { l.InlineMessages[0].ID = " " }, "'id'"},
		{"duplicate-id", func(l *Locale) { l.InlineMessages = append(l.InlineMessages, l.InlineMessages[0]) }, "duplicate"},
		{"empty-selector", func(l *Locale) { l.InlineMessages[0].Selector = " " }, "'selector'"},
		{"empty-source", func(l *Locale) { l.InlineMessages[0].Source = "" }, "'source'"},
		{"empty-template", func(l *Locale) { l.InlineMessages[0].Template = "" }, "'template'"},
		{"no-slots", func(l *Locale) { l.InlineMessages[0].Slots = nil }, "at least one slot"},
		{"slot-with-no-target", func(l *Locale) { l.InlineMessages[0].Slots[0] = InlineSlot{} }, "exactly one"},
		{"slot-with-two-targets", func(l *Locale) {
			l.InlineMessages[0].Slots[0].TextIndex = l.InlineMessages[0].Slots[1].TextIndex
		}, "exactly one"},
		{"blank-slot-selector", func(l *Locale) { l.InlineMessages[0].Slots[0].Selector = " " }, "non-empty 'selector'"},
		{"negative-text-index", func(l *Locale) { *l.InlineMessages[0].Slots[1].TextIndex = -1 }, "negative"},
		{"text-label-on-text-node", func(l *Locale) { l.InlineMessages[0].Slots[1].Text = "Account" }, "only with 'selector'"},
		{"missing-source-slot", func(l *Locale) { l.InlineMessages[0].Source = "Open {1}." }, "'source'"},
		{"reordered-source-slots", func(l *Locale) { l.InlineMessages[0].Source = "Open {2} for {1}." }, "'source'"},
		{"duplicated-source-slot", func(l *Locale) { l.InlineMessages[0].Source += " {2}" }, "'source'"},
		{"extra-source-brace", func(l *Locale) { l.InlineMessages[0].Source += " {name}" }, "'source'"},
		{"missing-template-slot", func(l *Locale) { l.InlineMessages[0].Template = "{1} を開く。" }, "'template'"},
		{"reordered-template-slots", func(l *Locale) { l.InlineMessages[0].Template = "{2} のために {1} を開く。" }, "'template'"},
		{"duplicated-template-slot", func(l *Locale) { l.InlineMessages[0].Template += " {2}" }, "'template'"},
		{"extra-template-slot", func(l *Locale) { l.InlineMessages[0].Template += " {3}" }, "'template'"},
		{"malformed-template-slot", func(l *Locale) { l.InlineMessages[0].Template = "{01} を {2} のために開く。" }, "'template'"},
		{"unmatched-template-brace", func(l *Locale) { l.InlineMessages[0].Template += "}" }, "'template'"},
		{"no-node-for-prefix", func(l *Locale) {
			l.InlineMessages[0].Source = "{1} for {2}."
			l.InlineMessages[0].Template = "開く {1} for {2}."
		}, "cannot add a prefix"},
		{"no-node-for-suffix", func(l *Locale) {
			l.InlineMessages[0].Source = "Open {1} for {2}"
			l.InlineMessages[0].Template = "{1} を {2} のために開く。"
		}, "cannot add a suffix"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			loc := inlineLocaleFixture()
			tc.edit(loc)
			if err := validateLocale("ja", loc); err == nil || !strings.Contains(err.Error(), tc.want) || !strings.Contains(err.Error(), "inline message") {
				t.Fatalf("expected inline-message error containing %q, got %v", tc.want, err)
			}
		})
	}
	for _, tc := range []struct{ source, template string }{
		{"Open {1} for {2}.", "{1} を {2} のために開く。"},
		{"{1} was sent to {2}", "{1} の送信先は {2}"},
		{"Read {1} and {2}", "参照先：{1} と {2}"},
	} {
		loc := inlineLocaleFixture()
		loc.InlineMessages[0].Source, loc.InlineMessages[0].Template = tc.source, tc.template
		if err := validateLocale("ja", loc); err != nil {
			t.Errorf("valid inline message rejected: %v", err)
		}
	}
}

// textIndex zero is a real child-node position, not an omitted target. Check
// the actual preload serializer so inline metadata reaches the DOM engine.
func TestPreloadIncludesInlineMessages(t *testing.T) {
	zero := 0
	loc := inlineLocaleFixture()
	loc.InlineMessages[0].Slots[1].TextIndex = &zero
	store := &LocaleStore{engineFragment: []byte("const AG_LOCALE = LOCALE_DATA_PLACEHOLDER;")}
	fragment, err := store.buildPreloadFragment(loc)
	if err != nil {
		t.Fatal(err)
	}
	payload := strings.TrimSuffix(strings.TrimPrefix(fragment, "const AG_LOCALE = "), ";")
	var decoded struct {
		InlineMessages []InlineMessage `json:"inlineMessages"`
	}
	if err := json.Unmarshal([]byte(payload), &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded.InlineMessages, loc.InlineMessages) {
		t.Fatalf("inline metadata was lost in preload: got %+v, want %+v", decoded.InlineMessages, loc.InlineMessages)
	}
	if !strings.Contains(payload, `"textIndex":0`) {
		t.Fatal("preload omitted child-node position zero")
	}
}
