package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Tasker is the interface to process
type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func buildFilesList(srcDir string) []string {
	list := []string{}
	fmt.Println("Generating file list ...")
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, "jpg") {
			return nil
		}
		list = append(list, path)
		return nil
	})
	return list
}
