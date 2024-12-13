package parameters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/maurotory/jenkins-cli/pkg/errors"
)

func GetParameters(configPath string) (map[string]string, error) {
	params := make(map[string]string)

	file, err := os.Open("params.json")
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.OpenFile, err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.ReadFile, err)
	}

	err = json.Unmarshal(bytes, &params)
	if err != nil {
		return nil, fmt.Errorf("%s, %v\n", errors.ParseJson, err)
	}

	return params, nil
}
