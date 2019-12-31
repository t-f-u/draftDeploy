package read

import (
	"path/filepath"
	"testing"
)

const TestDraftDir = "../../test_draft"

func TestReadFiles(t *testing.T) {
	abs, _ := filepath.Abs(TestDraftDir)
	t.Log(abs)
	files, err := readFiles(abs)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fail()
	}
	for i, file := range files {
		t.Log(i, "Name", file.Name)
		t.Log(i, "Ext", file.Ext)
		t.Log(i, "Path", file.Path)
		t.Log(i, "FullPath", file.FullPath)
	}
}

func TestIsNotEmptyAll(t *testing.T) {
	if !(isNotEmptyAll("a", "b", "c")) {
		t.Fail()
	}
	if isNotEmptyAll("", "b", "") {
		t.Fail()
	}
	var str string
	if isNotEmptyAll(str) {
		t.Fail()
	}
}

func TestGetBuildConfig(t *testing.T) {
	abs, _ := filepath.Abs(TestDraftDir)
	t.Log(abs)
	files, err := readFiles(abs)
	if err != nil {
		t.Fatal(err)
	}
	bc, err := getBuildConfig(files)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(bc.Title)
		t.Log(bc.Chapters)
	}
}

func TestGetDrafts(t *testing.T) {
	abs, _ := filepath.Abs(TestDraftDir)
	t.Log(abs)
	files, err := readFiles(abs)
	if err != nil {
		t.Fatal(err)
	}
	bc, err := getBuildConfig(files)
	if err != nil {
		t.Fatal(err)
	}
	dfs, err := getDraftFiles(bc, files)
	if err != nil {
		t.Fatal(err)
	}
	ds, err := readDrafts(dfs)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(ds)
	}
}
