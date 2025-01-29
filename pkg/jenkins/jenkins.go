package jenkins

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bndr/gojenkins"
	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/errors"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Blue = "\033[34m"
var Gray = "\033[37m"

var folderType = "com.cloudbees.hudson.plugins.folder.Folder"
var jobType = "org.jenkinsci.plugins.workflow.job.WorkflowJob"
var freestyleType = "hudson.model.FreeStyleProject"
var organizationType = "jenkins.branch.OrganizationFolder"

var timerType = "org.jenkinsci.plugins.parameterizedscheduler.ParameterizedTimerTriggerCause"
var upstreamType = "org.jenkinsci.plugins.workflow.support.steps.build.BuildUpstreamCause"

type JenkinsClient struct {
	client *gojenkins.Jenkins
	ctx    context.Context
}

func ConnectToJenkins(conf *config.JenkinsConfig) (*JenkinsClient, error) {
	jenkins := gojenkins.CreateJenkins(nil, conf.Host, conf.User, conf.Token)
	ctx := context.Background()
	_, err := jenkins.Init(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.JenkinsConf, err)
	}
	return &JenkinsClient{client: jenkins, ctx: ctx}, nil
}

func (j JenkinsClient) Info() {
	fmt.Println("Successfully Connected:")
	fmt.Printf("Server: %s\n", j.client.Server)
	fmt.Printf("Version: %s\n", j.client.Version)
	// fmt.Printf("Raw: %s\n", j.client.Raw.Jobs)
}

func (j JenkinsClient) ListBuilds(jobId string, maxQuantity int) (*gojenkins.JobBuild, error) {
	jobId, err := parseJobId(jobId)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
	}

	builds, err := j.client.GetAllBuildIds(j.ctx, jobId)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.GetBuilds, err)
	}
	printColumnInfo([]string{"Build ID", "User", "Start Time", "Duration", "Result"}, 15)
	count := 0
	for _, build := range builds {
		if count >= maxQuantity {
			break
		}
		buildId := build.Number
		data, err := j.client.GetBuild(j.ctx, jobId, buildId)
		if err != nil {
			return nil, fmt.Errorf("%s, %v\n", errors.GetBuild, err)
		}
		user, err := getUser(data)
		if err != nil {
			return nil, fmt.Errorf("%s, %v\n", errors.GetUser, err)
		}

		var result string
		if data.IsRunning(j.ctx) {
			result = "RUNNING"
		} else if data.GetResult() == "SUCCESS" {
			result = data.GetResult()
		} else if data.GetResult() == "FAILURE" {
			result = data.GetResult()
		} else if data.GetResult() == "ABORTED" {
			result = data.GetResult()
		} else {
			return nil, fmt.Errorf("%s, %v\n", errors.WrongJobResult, err)
		}

		duration := fmt.Sprintf("%s", time.Duration(data.GetDuration()*float64(time.Millisecond)))
		timestamp := data.GetTimestamp()
		startDate := fmt.Sprintf("%02d-%02d %02d:%02d",
			timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute())
		printColumnInfo([]string{fmt.Sprintf("%d", buildId), user, startDate,
			duration, result}, 15)
		count++
	}

	return nil, nil
}

func (j JenkinsClient) ListArtifacts(jobId string, buildId int64, latest bool) error {
	jobId, err := parseJobId(jobId)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
	}
	if latest {
		builds, err := j.client.GetAllBuildIds(j.ctx, jobId)
		if err != nil {
			return fmt.Errorf("%s, %v\n", errors.GetBuilds, err)
		}
		buildId = builds[0].Number
	}
	build, err := j.client.GetBuild(j.ctx, jobId, buildId)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.GetBuild, err)
	}
	artifacts := build.GetArtifacts()
	for _, artifact := range artifacts {
		fmt.Println(artifact.FileName)
	}

	return nil
}
func (j JenkinsClient) ListJobs(folderId string, view string, maxQuantity int) (*gojenkins.JobBuild, error) {
	if folderId == "" {
		views, err := j.client.GetAllViews(j.ctx)
		if err != nil {
			return nil, fmt.Errorf("%s, %v\n", errors.GetView, err)
		}
		if view != "" {
			for _, v := range views {
				if view == v.GetName() {
					err := printJobs(v.Raw.Jobs, maxQuantity)
					if err != nil {
						return nil, fmt.Errorf("%v", err)
					}
				}
			}
		} else {
			for _, v := range views {
				fmt.Println(v.GetName())
				err := printJobs(v.Raw.Jobs, maxQuantity)
				if err != nil {
					return nil, fmt.Errorf("%v", err)
				}
				fmt.Println()
			}
		}
	} else {
		folderId, err := parseJobId(folderId)
		if err != nil {
			return nil, fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
		}

		folder, err := j.client.GetFolder(j.ctx, folderId)
		if err != nil {
			return nil, fmt.Errorf("%s, %v\n", errors.GetFolder, err)
		}
		err = printJobs(folder.Raw.Jobs, maxQuantity)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}
	return nil, nil
}

func (j JenkinsClient) ListViews(maxQuantity int) error {
	views, err := j.client.GetAllViews(j.ctx)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.GetView, err)
	}
	for _, view := range views {
		fmt.Println(view.GetName())
	}
	return nil
}

func (j JenkinsClient) GetBuild(jobId string, buildId int64, latest bool) error {
	jobId, err := parseJobId(jobId)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
	}
	if latest {
		builds, err := j.client.GetAllBuildIds(j.ctx, jobId)
		if err != nil {
			return fmt.Errorf("%s, %v\n", errors.GetBuilds, err)
		}
		buildId = builds[0].Number
	}
	build, err := j.client.GetBuild(j.ctx, jobId, buildId)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.GetBuild, err)
	}
	fmt.Println(build.GetResult())
	var result string
	if build.IsRunning(j.ctx) {
		result = Blue + "RUNNING" + Reset
	} else if build.GetResult() == "SUCCESS" {
		result = Green + build.GetResult() + Reset
	} else if build.GetResult() == "FAILURE" {
		result = Red + build.GetResult() + Reset
	} else if build.GetResult() == "ABORTED" {
		result = Gray + build.GetResult() + Reset
	} else {
		return fmt.Errorf("%s, %v\n", errors.WrongJobResult, err)
	}

	user, err := getUser(build)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.GetUser, err)
	}
	duration := time.Duration(build.GetDuration() * float64(time.Millisecond)).String()
	timestamp := build.GetTimestamp()
	startDate := fmt.Sprintf("%02d-%02d-%04d %02d:%02d:%02d",
		timestamp.Day(), timestamp.Month(), timestamp.Year(), timestamp.Hour(), timestamp.Minute(), timestamp.Second())
	separation := "----------------\n"
	fmt.Printf("Result:\n%s\n%s", result, separation)
	fmt.Printf("User:\n%s\n%s", user, separation)
	fmt.Printf("Duration:\n%s\n%s", duration, separation)
	fmt.Printf("Start Date:\n%s\n%s", startDate, separation)
	fmt.Printf("Artifacts:\n")
	artifacts := build.GetArtifacts()
	for _, artifact := range artifacts {
		fmt.Println(artifact.FileName)
	}
	fmt.Printf("%s", separation)
	fmt.Printf("Params:\n")
	params := build.GetParameters()
	for _, param := range params {
		fmt.Printf("%s: %s\n", param.Name, param.Value)
	}
	fmt.Printf("%s", separation)
	fmt.Printf("Upstream Jobs:\n")
	upstreamJobs, err := build.Job.GetUpstreamJobs(j.ctx)
	if err != nil {
		return err
	}
	for _, upJob := range upstreamJobs {
		fmt.Printf("%v\n", upJob.Raw)
	}

	fmt.Printf("%s", separation)
	upstreamNum, err := build.GetUpstreamBuildNumber(j.ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Upstream Build Number: %d\n", upstreamNum)

	fmt.Printf("%s", separation)
	fmt.Printf("Downstream jobs:\n")
	downstreamJobs, err := build.Job.GetDownstreamJobs(j.ctx)
	if err != nil {
		return err
	}
	for _, downJob := range downstreamJobs {
		fmt.Printf("%v\n", downJob)
	}

	fmt.Printf("%s", separation)
	fmt.Printf("Downstream Build Numbers:\n")
	downstreamNum, err := build.GetDownstreamBuilds(j.ctx)
	if err != nil {
		return err
	}
	for _, downNum := range downstreamNum {
		fmt.Printf("%v\n", downNum.Raw.Number)
	}
	return nil
}

func (j JenkinsClient) GetArtifact(jobId, artifact, output string, buildId int64, latest bool, printArtifact bool) error {
	jobId, err := parseJobId(jobId)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
	}
	if output == "" {
		currentFolder, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		output = currentFolder + "/" + artifact
	}
	if latest {
		builds, err := j.client.GetAllBuildIds(j.ctx, jobId)
		if err != nil {
			return fmt.Errorf("%s, %v\n", errors.GetBuilds, err)
		}
		buildId = builds[0].Number
	}
	build, err := j.client.GetBuild(j.ctx, jobId, buildId)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.GetBuild, err)
	}
	artifacts := build.GetArtifacts()
	for _, a := range artifacts {
		if a.FileName == artifact {
			if printArtifact {
				data, err := a.GetData(j.ctx)
				if err != nil {
					return err
				}
				if !isTextFile(data) {
					fmt.Printf("%s: Is a binary file.\n", a.FileName)
					return nil
				}
				fmt.Printf("%s", data)
				return nil

			}
			saveFile, err := checkFile(output)
			if err != nil {
				return fmt.Errorf("%v\n", err)
			}
			if !saveFile {
				fmt.Println("Aborting...")
				return nil
			}

			saved, err := a.Save(j.ctx, output)
			if err != nil {
				return fmt.Errorf("%s, %v\n", errors.SaveFile, err)
			}
			if !saved {
				return fmt.Errorf("%s\n", errors.SaveFile)
			}
			fmt.Printf("Artifact saved in the path: %s\n", output)
			return nil
		}
	}
	return fmt.Errorf("%s\n", errors.ArtifactNotFound)
}

func (j JenkinsClient) CreateJob(jobId string, params map[string]string) (int64, error) {
	jobId, err := parseJobId(jobId)
	if err != nil {
		return 0, fmt.Errorf("%s, %v\n", errors.ParseJobId, err)
	}
	queueid, err := j.client.BuildJob(j.ctx, jobId, params)
	if err != nil {
		return 0, fmt.Errorf("%s, %v\n", errors.CreateJob, err)
	}

	var item *gojenkins.Task
	timeout := 50
	var count int
	for count < timeout {
		item, err = j.client.GetQueueItem(j.ctx, queueid)
		if err != nil {
			return 0, fmt.Errorf("%s, %v\n", errors.GetbuildFromQueue, err)
		}
		if item.Raw.Executable.Number != 0 {
			break
		}
		time.Sleep(3 * time.Second)
		count++
	}

	return item.Raw.Executable.Number, nil
}

func (j JenkinsClient) Logs(jobId string, buildId int64, follow bool, latest bool) (*gojenkins.JobBuild, error) {
	jobId, err := parseJobId(jobId)
	if latest {
		builds, err := j.client.GetAllBuildIds(j.ctx, jobId)
		if err != nil {
			return nil, fmt.Errorf("%s, %v\n", errors.GetBuilds, err)
		}
		buildId = builds[0].Number
	}
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
				fmt.Printf("%s", console.Content)
			}

			time.Sleep(3 * time.Second)

			start = console.Offset
			requestLogs = console.HasMoreText
		}
	}

	return nil, nil
}
