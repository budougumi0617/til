package zip

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestUnzip(t *testing.T) {
	// do nothing.
}

// Unzip unzips files.
func Unzip(src, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	// ループの中でdeferをしてはいけない。
	deflate := func(f *zip.File) error {
		origin, err := f.Open()
		if err != nil {
			log.Printf("unzip: open %q", err)
			return err
		}
		defer origin.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			log.Printf("unzip: create dir %q", path)
			os.MkdirAll(path, f.Mode())
		} else {
			b, err := ioutil.ReadAll(origin)
			if err != nil {
				log.Printf("unzip: readall: %q", err)
				return err
			}

			// os.OpenFileは中間ディレクトリまでは作成しないので事前に作成する。
			dirs := filepath.Dir(f.Name)
			dirpath := filepath.Join(dest, dirs)
			os.MkdirAll(dirpath, 0755)

			if err := ioutil.WriteFile(path, b, f.Mode()); err != nil {
				log.Printf("unzip: writefile: %q", err)
				return err
			}
		}
		return nil
	}

	for _, f := range reader.File {
		if err := deflate(f); err != nil {
			return err
		}
	}
	return nil
}
