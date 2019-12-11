package read

import (
	"path/filepath"
	"testing"
)

func TestReadFiles(t *testing.T) {
	abs, _ := filepath.Abs("../test_draft")
	t.Log(abs)
	files, err := readFiles(abs)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fail()
	}
	for _, file := range files {
		t.Log(file.Name)
		t.Log(file.Ext)
		t.Log(file.Path)
		t.Log(file.FullPath)
	}
}
