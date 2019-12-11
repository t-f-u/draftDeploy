package read

import (
	"bufio"
	. "draftDeploy/app/model"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func Read(args Args) (ReadData, error) {
	var readData ReadData
	files, err := readFiles(args.InputDir)
	if err != nil {
		return readData, nil
	}
	bc, err := getBuildConfig(files)
	if err != nil {
		return readData, err
	}
	draftFiles, err := getDraftFiles(bc, files)
	if err != nil {
		return readData, err
	}
	drafts, err := readDrafts(draftFiles)
	if err != nil {
		return readData, err
	}
	readData.BuildConfig = bc
	readData.Drafts = drafts
	return readData, nil
}

func readFiles(dir string) ([]File, error) {
	var ret []File
	err := filepath.Walk(dir, func(fullPath string, info os.FileInfo, err error) error {
		ext := filepath.Ext(fullPath)
		rep := regexp.MustCompile(ext + "$")
		path, name := filepath.Split(rep.ReplaceAllString(fullPath, ""))
		if isNotEmptyAll(name, ext, path, fullPath) {
			ret = append(ret, File{
				Name:     name,
				Ext:      ext,
				Path:     path,
				FullPath: fullPath,
			})
		}
		return err
	})
	return ret, err
}

func isNotEmptyAll(strings ...string) bool {
	for _, str := range strings {
		if str == "" {
			return false
		}
	}
	return true
}

func getBuildConfig(files []File) (BuildConfig, error) {
	var bc BuildConfig
	buildYml, err := findBuildYaml(files)
	if err != nil {
		return bc, err
	}
	data, err := ioutil.ReadFile(buildYml.FullPath)
	if err != nil {
		return bc, err
	}
	err = yaml.Unmarshal(data, &bc)
	return bc, err
}

func findBuildYaml(files []File) (File, error) {
	for _, file := range files {
		if file.Name == "build" && isYaml(file) {
			return file, nil
		}
	}
	return File{}, errors.New("build.yml is not found")
}

func isYaml(file File) bool {
	return file.Ext == ".yaml" || file.Ext == ".yml"
}

func getDraftFiles(bc BuildConfig, files []File) ([]File, error) {
	ret := make([]File, len(bc.Chapters))
ForChapters:
	for i, ch := range bc.Chapters {
		for _, f := range files {
			if f.Name == ch && f.Ext == ".txt" {
				ret[i] = f
				continue ForChapters
			}
		}
	}
	if len(ret) == 0 {
		return ret, errors.New("no drafts")
	}
	return ret, nil
}

func readDrafts(draftFiles []File) ([]Draft, error) {
	var in *os.File
	defer func() {
		if in != nil {
			if err := in.Close(); err != nil {
				panic(err)
			}
		}
	}()
	var ret []Draft
	for _, file := range draftFiles {
		in, err := os.Open(file.FullPath)
		if err != nil {
			continue
		}
		var data []string
		sc := bufio.NewScanner(in)
		for sc.Scan() {
			data = append(data, sc.Text())
		}
		ret = append(ret, Draft{Data: data})
	}
	return ret, nil
}
