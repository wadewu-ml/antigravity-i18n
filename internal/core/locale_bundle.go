package core

import (
	"archive/zip"
	"compress/flate"
	"encoding/json"
	"io"
	"io/fs"
)

const localeIndexFile = "_index.json"

// WriteLocaleBundle validates the source dictionaries and packages each as a
// separately compressed ZIP entry. The index supplies the language list
// without decoding the dictionaries when the setup window opens.
func (s *LocaleStore) WriteLocaleBundle(out io.Writer) error {
	locales, err := s.describeLocaleFiles()
	if err != nil {
		return err
	}
	archive := zip.NewWriter(out)
	archive.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, flate.BestCompression)
	})
	for _, locale := range locales {
		name := s.localeDir + "/" + locale.Code + ".json"
		data, err := fs.ReadFile(s.localesFS, name)
		if err != nil {
			return err
		}
		entry, err := archive.Create(name)
		if err != nil {
			return err
		}
		if _, err := entry.Write(data); err != nil {
			return err
		}
	}
	index, err := archive.Create(s.localeDir + "/" + localeIndexFile)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(index).Encode(locales); err != nil {
		return err
	}
	return archive.Close()
}
