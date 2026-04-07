package loader

import (
	"os"
	"path/filepath"
)

func GetCSVFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".csv" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
