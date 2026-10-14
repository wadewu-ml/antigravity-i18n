package core

// Options controls an apply or restore run.
type Options struct {
	AppDir    string
	Locale    string
	Restart   bool
	Force     bool
	TimeoutMs int

	// Progress receives stage updates for a GUI.
	Progress func(stage, detail string)
}

// Service is the entry point the CLI and the GUI both drive.
type Service struct {
	store *LocaleStore
	// Progress receives stage updates for a GUI.
	Progress func(stage, detail string)
	// notes carries warnings that should be shown to the user but must not fail
	// an otherwise successful operation.
	notes []string
}

// NewService builds the service around the embedded locale pack.
func NewService(store *LocaleStore) *Service {
	return &Service{store: store}
}

// Notes returns warnings accumulated by the last operation.
func (s *Service) Notes() []string { return s.notes }

// DescribeLocales lists the bundled language packs with their coverage.
func (s *Service) DescribeLocales() ([]LocaleInfo, error) {
	return s.store.Describe()
}

// clearNotes drops warnings from previous operations.
func (s *Service) clearNotes() { s.notes = nil }

// progress reports a stage update when a sink is configured.
func (s *Service) progress(stage, detail string) {
	if s.Progress != nil {
		s.Progress(stage, detail)
	}
}
