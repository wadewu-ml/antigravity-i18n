package core

import "errors"

// ErrAntigravityRunning reports an installation that requires the user's choice.
var ErrAntigravityRunning = errors.New("Antigravity is running. Exit it and retry")

// ErrGracefulShutdownTimeout allows callers to offer an explicit force retry.
var ErrGracefulShutdownTimeout = errors.New("Antigravity did not exit before the timeout. Close it manually or explicitly choose forced shutdown and retry")

// Options controls an apply or restore run.
type Options struct {
	AppDir         string
	Locale         string
	Restart        bool
	Force          bool
	RequireClosed  bool // check for a running app before proceeding unless Force is set
	TimeoutMs      int
	EffortSlider   *bool // nil keeps the setting in the installed archive
	SessionMetrics *bool // nil keeps the setting in the installed archive
}

// EnglishLocale uses Antigravity's built-in language.
const EnglishLocale = "en"

// Service is the entry point the CLI and the GUI both drive.
type Service struct {
	store *LocaleStore
	// Progress receives stage updates for a GUI.
	Progress func(stage, detail string)
	// notes carries warnings that should be shown to the user but must not fail
	// an otherwise successful operation.
	notes []string
	// Process hooks keep fixture tests independent of the installed application.
	stopApp    func(*paths, bool, int) (bool, error)
	appRunning func(*paths) (bool, error)
	launchApp  func(*paths) error
}

// NewService builds the service around the embedded locale pack.
func NewService(store *LocaleStore) *Service {
	return &Service{store: store, stopApp: stopAntigravity, appRunning: isAntigravityRunning, launchApp: launchAntigravity}
}

// Notes returns warnings accumulated by the last operation.
func (s *Service) Notes() []string { return s.notes }

// DescribeLocales lists the bundled language packs with their coverage.
//
// English is provided by Antigravity, so it has no bundled translations.
func (s *Service) DescribeLocales() ([]LocaleInfo, error) {
	locales, err := s.store.Describe()
	if err != nil {
		return nil, err
	}
	return append([]LocaleInfo{{
		Code:         EnglishLocale,
		Name:         "English",
		EnglishName:  "English",
		Translations: 0,
		Coverage:     100,
	}}, locales...), nil
}

// clearNotes drops warnings from previous operations.
func (s *Service) clearNotes() { s.notes = nil }

func (s *Service) closeApp(p *paths, opts Options) (bool, error) {
	if opts.RequireClosed && !opts.Force {
		running, err := s.appRunning(p)
		if err != nil {
			return running, &OpError{Op: "checking Antigravity", Cause: err}
		}
		if running {
			return true, ErrAntigravityRunning
		}
		return false, nil
	}

	timeout := opts.TimeoutMs
	if timeout <= 0 {
		timeout = gracefulTimeoutMs
	}
	s.progress("正在关闭 Antigravity", p.AppDir)
	running, err := s.stopApp(p, opts.Force, timeout)
	if err != nil {
		return running, &OpError{Op: "closing Antigravity", Cause: err}
	}
	return running, nil
}

// progress reports a stage update when a sink is configured.
func (s *Service) progress(stage, detail string) {
	if s.Progress != nil {
		s.Progress(stage, detail)
	}
}
