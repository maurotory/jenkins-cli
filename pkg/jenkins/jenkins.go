package jenkins

import (
	"context"
	"fmt"
     "regexp"

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
    fmt.Printf("Server: %s\n",j.client.Server)
    fmt.Printf("Version: %s\n", j.client.Version)
    fmt.Printf("Raw: %s\n", j.client.Raw.Jobs)
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
    fmt.Printf("|   jobID   |    Result    |\n")
    for _, build := range builds {
        buildId := build.Number
        data, err := j.client.GetBuild(j.ctx, jobId, buildId)
        if err != nil {
              return nil, fmt.Errorf("%s, %v\n", errors.CompileRegex, err)
        }
        var result string
        if data.IsRunning(j.ctx) {
            result = Blue + "RUNNING" + Reset
        } else if data.GetResult() == "SUCCESS" {
        result =Green +  data.GetResult() + Reset
        } else if data.GetResult() == "FAILURE" {
            result =Red +  data.GetResult() + Reset
        } else {
              return nil, fmt.Errorf("%s, %v\n", errors.WrongJobResult, err)
        }
        fmt.Printf("|     %d    |    %s    |\n", buildId, result)
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
