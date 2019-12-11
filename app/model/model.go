package model

type Args struct {
	InputDir  string
	OutputDir string
}

type File struct {
	Name     string
	Ext      string
	Path     string
	FullPath string
}
