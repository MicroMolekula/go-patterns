package composite

import (
	"os"
	"path/filepath"
)

func BuildTree(path string) (Component, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return &File{name: fileInfo.Name(), size: int(fileInfo.Size())}, nil
	}

	folder := NewFolder(fileInfo.Name(), nil)
	childs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, child := range childs {
		fullPath := filepath.Join(path, child.Name())
		c, err := BuildTree(fullPath)
		if err != nil {
			return nil, err
		}
		folder.Add(c)
	}
	return folder, nil
}
