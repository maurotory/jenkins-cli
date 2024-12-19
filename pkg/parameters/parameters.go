package parameters

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/maurotory/jenkins-cli/pkg/errors"
)

func GetParameters(configPath string) (map[string]string, error) {
	params, err := godotenv.Read(configPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errors.ParseParamsFile, err)
	}
	return params, nil
}
