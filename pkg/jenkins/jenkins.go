package jenkins

import (
	"context"
	"fmt"

	"github.com/bndr/gojenkins"
	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/errors"
)

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

func (j JenkinsClient) Whoami() error {
	resp, err := j.client.Info(j.ctx)
	if err != nil {
		return fmt.Errorf("%s, %v\n", errors.Whoami, err)
	}
	fmt.Println(resp)
	return nil
}
