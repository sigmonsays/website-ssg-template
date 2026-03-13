package site

import (
	"io/fs"
	"os"
	"path/filepath"
)

type FileData struct {
	Name    string
	Path    string
	Content string
}

func ListFiles(dir string) ([]FileData, error) {
	var files []FileData
	err := filepath.Walk(dir, func(path string, d fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			relpath, _ := filepath.Rel(dir, path)
			files = append(files, FileData{
				Name:    relpath,
				Path:    path,
				Content: string(data),
			})
		}
		return nil
	})
	return files, err
}
