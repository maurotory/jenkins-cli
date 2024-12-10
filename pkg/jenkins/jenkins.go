package jenkins

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/bndr/gojenkins"
	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/errors"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type JenkinsClient struct {
	client *gojenkins.Jenkins
	ctx    context.Context
}

func ConnectToJenkins(conf *config.JenkinsConfig) (*JenkinsClient, error) {
	jenkins := gojenkins.CreateJenkins(nil, conf.Host, conf.User, conf.Password)
	ctx := context.Background()
	_, err := jenkins.Init(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.JenkinsConf, err)
	}
	return &JenkinsClient{client: jenkins, ctx: ctx}, nil
}

func (j JenkinsClient) Info() {
	fmt.Printf("Server: %s\n", j.client.Server)
	fmt.Printf("Version: %s\n", j.client.Version)
	// fmt.Printf("Raw: %s\n", j.client.Raw.Jobs)
}

func (j JenkinsClient) ListBuilds(jobId string) (*gojenkins.JobBuild, error) {
	jobId, err := parseJobId(jobId)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
	}

	builds, err := j.client.GetAllBuildIds(j.ctx, jobId)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.GetBuilds, err)
	}

	printBuildInfo([]string{"JobID", "User", "Result"})

	for _, build := range builds {
		buildId := build.Number
		data, err := j.client.GetBuild(j.ctx, jobId, buildId)
		if err != nil {
			return nil, fmt.Errorf("%s, %v\n", errors.GetBuild, err)
		}
		user := data.Raw.Actions[0].Causes[0]["userId"].(string)

		var result string
		if data.IsRunning(j.ctx) {
			result = "RUNNING"
		} else if data.GetResult() == "SUCCESS" {
			result = data.GetResult()
		} else if data.GetResult() == "FAILURE" {
			result = data.GetResult()
		} else {
			return nil, fmt.Errorf("%s, %v\n", errors.WrongJobResult, err)
		}
		printBuildInfo([]string{fmt.Sprintf("%d", buildId), user, result})
	}

	return nil, nil
}

func (j JenkinsClient) ListItems(jobId string) (*gojenkins.JobBuild, error) {
	return nil, nil
}

func (j JenkinsClient) CreateJob(jobId string) (*gojenkins.JobBuild, error) {
	return nil, nil
}

func (j JenkinsClient) Logs(jobId string, buildId int64, follow bool) (*gojenkins.JobBuild, error) {
	jobId, err := parseJobId(jobId)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
	}
	build, err := j.client.GetBuild(j.ctx, jobId, buildId)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.GetBuild, err)
	}
	if !follow {
		fmt.Println(build.GetConsoleOutput(j.ctx))
	} else {
		var start int64 = 0
		requestLogs := true
		for requestLogs {
			console, err := build.GetConsoleOutputFromIndex(j.ctx, start)
			if err != nil {
				return nil, fmt.Errorf("%s, %v\n", errors.ConsoleOutput, err)
			}

			if console.Offset != start {
				fmt.Printf(console.Content)
			}

			time.Sleep(3 * time.Second)

			start = console.Offset
			requestLogs = console.HasMoreText
		}
	}

	return nil, nil
}

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

func printBuildInfo(info []string) {
	columnSize := 10
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
