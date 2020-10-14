package tar

import (
	"../../areader"
	"archive/tar"
	"io"
	"os"
)

func filenames(path string) []string {
	r, _ := os.Open(path)
	defer r.Close()

	var names []string
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		names = append(names, hdr.Name)
	}
	return names
}

func init() {
	areader.RegisterFormat("tar", filenames)
}
