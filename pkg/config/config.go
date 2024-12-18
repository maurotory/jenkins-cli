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
var tokenVar = "token"

type JenkinsConfig struct {
	Host  string
	User  string
	Token string
}

func GetConfig(configPath string) (*JenkinsConfig, error) {
	var host, user, token string
	if configPath == "" {
		homePath := os.Getenv(homePathVar)
		if homePath == "" {
			return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, homePath)
		}
		viper.SetConfigName("config.json")
		viper.AddConfigPath(homePath + "/.jctl")
		viper.SetConfigType("json")
	} else {
		viper.SetConfigFile(configPath)
		viper.SetConfigType("json")
	}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	host = viper.GetString(hostVar)
	if host == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, hostVar)
	}
	user = viper.GetString(userVar)
	if user == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, userVar)
	}
	token = viper.GetString(tokenVar)
	if token == "" {
		return nil, fmt.Errorf("%s: %s\n", errors.EmptyVar, tokenVar)
	}

	return &JenkinsConfig{
		Host:  host,
		User:  user,
		Token: token,
	}, nil
}
