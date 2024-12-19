package jenkins

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bndr/gojenkins"
	"github.com/maurotory/jenkins-cli/pkg/errors"
)

func parseJobId(jobId string) (string, error) {
	r, err := regexp.Compile(`\/$`)
	if err != nil {
		return "", fmt.Errorf("%s, %v\n", errors.CompileRegex, err)
	}
	jobId = r.ReplaceAllString(jobId, "")

	r, err = regexp.Compile(`^\/`)
	if err != nil {
		return "", fmt.Errorf("%s, %v\n", errors.CompileRegex, err)
	}
	jobId = r.ReplaceAllString(jobId, "")

	r, err = regexp.Compile(`\/`)
	if err != nil {
		return "", fmt.Errorf("%s, %v\n", errors.CompileRegex, err)
	}
	jobId = r.ReplaceAllString(jobId, "/job/")

	return jobId, nil
}

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

func checkFile(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("%s", err)
	}
	if os.IsNotExist(err) {
		return true, nil
	}
	if info.IsDir() {
		return false, fmt.Errorf("%s: %s", errors.PathIsADirectory, filePath)
	}
	reader := bufio.NewReader(os.Stdin)
	s := fmt.Sprintf("File: %s already exists, are you sure you want to override it?", filePath)
	fmt.Printf("%s [y/n]: ", s)

	resp, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("%v", err)
	}

	resp = strings.ToLower(strings.TrimSpace(resp))
	if resp == "y" || resp == "yes" {
		return true, nil
	} else {
		return false, nil
	}
}
