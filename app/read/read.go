package read

import (
	. "draftDeploy/app/model"
	"os"
	"path/filepath"
	"regexp"
)

func Read(args Args) (string, error) {

	return "", nil
}

func readFiles(dir string) ([]File, error) {
	var ret []File
	err := filepath.Walk(dir, func(fullPath string, info os.FileInfo, err error) error {
		ext := filepath.Ext(fullPath)
		rep := regexp.MustCompile(ext + "$")
		path, name := filepath.Split(rep.ReplaceAllString(fullPath, ""))
		ret = append(ret, File{
			Name:     name,
			Ext:      ext,
			Path:     path,
			FullPath: fullPath,
		})
		return err
	})
	return ret, err
}
