package app

func Start(dir string, dist string) error {
	args, err := validateArgs(dir, dist)
	if err != nil {
		return err
	}
	return write(convert(read(args)))
}
