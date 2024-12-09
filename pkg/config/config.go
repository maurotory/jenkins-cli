package config

import (
	"fmt"
	"os"

	"github.com/maurotory/jenkins-cli/pkg/errors"
	"github.com/spf13/viper"
)

var homePathVar = "HOME"

var hostVar = "host"
var userVar = "user"
var passwordVar = "password"

type JenkinsConfig struct {
	Host  string
	User  string
	Password string
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
	host := viper.GetString(hostVar)
	if host == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, hostVar)
	}
	user := viper.GetString(userVar)
	if user == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, userVar)
	}
	password := viper.GetString(passwordVar)
	if password == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, passwordVar)
	}

	return &JenkinsConfig{
		Host:  host,
		User:  user,
		Password: password,
	}, nil
}
