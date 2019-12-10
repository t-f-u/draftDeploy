package app

func validateArgs(dir string, dist string) (Args, error) {
	return Args{
		InputDir:  dir,
		OutputDir: dist,
	}, nil
}
