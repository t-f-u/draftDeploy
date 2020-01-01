package app

import (
	. "draftDeploy/app/convert"
	. "draftDeploy/app/read"
	. "draftDeploy/app/write"
	"log"
	"os"
)

func Start(dir string, dist string) error {
	args, err := validateArgs(dir, dist)
	if err != nil {
		return err
	}
	return Write(Convert(Read(args)))
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("there must be 2 arguments: target and distribution")
	}

	if err := Start(os.Args[1], os.Args[2]); err != nil {
		log.Fatal(err)
	}
}
