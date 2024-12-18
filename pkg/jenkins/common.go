package jenkins

import (
	"fmt"

	"github.com/bndr/gojenkins"
	"github.com/maurotory/jenkins-cli/pkg/errors"
)

func printJobs(jobs []gojenkins.InnerJob) error {
	printColumnInfo([]string{"Name", "Type"}, 15)

	for _, job := range jobs {
		var itemType string
		if job.Class == jobType {
			itemType = "Pipeline"
		} else if job.Class == folderType {
			itemType = "Folder"
		} else {
			return fmt.Errorf("%s\n", errors.UnknownItemType)
		}

		printColumnInfo([]string{job.Name, itemType}, 15)
	}
	return nil
}

func printColumnInfo(info []string, columnSize int) {
	row := "|"
	for _, param := range info {
		paramLength := len(param)
		if param == "RUNNING" {
			param = Blue + param + Reset
		} else if param == "SUCCESS" {
			param = Green + param + Reset
		} else if param == "FAILURE" {
			param = Red + param + Reset
		}
		row = row + param
		for i := 0; i < columnSize-paramLength; i++ {
			row = row + " "
		}
		row = row + "|"
	}
	fmt.Printf("%s\n", row)
}
