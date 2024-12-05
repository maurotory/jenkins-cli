package config

import (
	"fmt"
	"os"

	"github.com/maurotory/jenkins-cli/pkg/errors"
	"github.com/spf13/viper"
)

var homePathVar = "HOME"

type JenkinsConfig struct {
	Host  string
	User  string
	Token string
}

func GetConfig() (*JenkinsConfig, error) {
	homePath := os.Getenv(homePathVar)
	if homePath == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, homePath)
	}
	viper.SetConfigName("config.json")
	viper.AddConfigPath(homePath + "/.jctl")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	host := viper.GetString("host")
	if host == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, host)
	}
	user := viper.GetString("user")
	if host == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, user)
	}
	token := viper.GetString("token")
	if host == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, token)
	}

	return &JenkinsConfig{
		Host:  host,
		User:  user,
		Token: token,
	}, nil
}
