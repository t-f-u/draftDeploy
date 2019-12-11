package app

import (
	. "draftDeploy/app/model"
	"errors"
	"os"
	"path/filepath"
)

func validateArgs(dir string, dist string) (Args, error) {
	emptyArgs := Args{}
	if !(filepath.IsAbs(dir) && filepath.IsAbs(dist)) {
		return emptyArgs, errors.New("dir and dist must be absolute path")
	}
	if !(isDirectory(dir) && isDirectory(dist)) {
		return emptyArgs, errors.New("dir and dist do not exist or are not directories")
	}
	return Args{
		InputDir:  dir,
		OutputDir: dist,
	}, nil
}

func isDirectory(path string) bool {
	if f, err := os.Stat(path); err == nil && f.IsDir() {
		return true
	} else {
		return false
	}
}
