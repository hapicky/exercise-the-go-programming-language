package zip

import (
	"../../areader"
	"archive/zip"
)

func filenames(path string) []string {
	r, _ := zip.OpenReader(path)
	defer r.Close()

	names := make([]string, len(r.File))
	for i, f := range r.File {
		names[i] = f.Name
	}
	return names
}

func init() {
	areader.RegisterFormat("zip", filenames)
}
