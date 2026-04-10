package loader

import (
	"os"
	"path/filepath"
)

func GetCSVFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".csv" {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	return files, nil
}
