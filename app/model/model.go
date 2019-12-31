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

func (d Draft) EachLines(f func(string) string) Draft {
	var ret []string
	for _, str := range d.Data {
		ret = append(ret, f(str))
	}
	return Draft{Data: ret}
}

type ReadData struct {
	OutputDir   string
	BuildConfig BuildConfig
	Drafts      []Draft
}

type ConvertData struct {
	OutputDir   string
	BuildConfig BuildConfig
	Data        []string
}
