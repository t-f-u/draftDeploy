package app

import (
	. "draftDeploy/app/convert"
	. "draftDeploy/app/read"
	. "draftDeploy/app/write"
)

func Start(dir string, dist string) error {
	args, err := validateArgs(dir, dist)
	if err != nil {
		return err
	}
	return Write(Convert(Read(args)))
}
