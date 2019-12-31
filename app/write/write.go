package write

import (
	. "draftDeploy/app/model"
	"fmt"
	"os"
)

func Write(data ConvertData, err error) error {
	if err != nil {
		return err
	}
	distFile := data.OutputDir + "/" + data.BuildConfig.Title + ".txt"
	file, err := os.Create(distFile)
	if err != nil {
		return err
	}
	defer file.Close()
	sum := 0
	for _, line := range data.Data {
		n, err := file.WriteString(line)
		if err != nil {
			return err
		}
		sum += n
	}
	fmt.Print(sum)
	return nil
}
