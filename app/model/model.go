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

type BuildConfig struct {
	Title    string
	Chapters []string
}

type Draft struct {
	Data []string
}

type ReadData struct {
	BuildConfig BuildConfig
	Drafts      []Draft
}

type ConvertData struct {
	Data []byte
}
