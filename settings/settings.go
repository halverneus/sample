package settings

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	// Get any settings value.
	Get struct {

		// Database connection parameters.
		Database struct {
			Host           string `yaml:"host"`
			Port           int    `yaml:"port"`
			Name           string `yaml:"name"`
			User           string `yaml:"user"`
			Password       string `yaml:"password"`
			MaxConnections int    `yaml:"max-connections"`
		} `yaml:"database"`

		// Storage for files.
		Storage struct {
			Folder string `yaml:"folder"`
		} `yaml:"storage"`

		BindAddress string `yaml:"bind"`
	}
)

// Parse the settings file.
func Parse(filename string) (err error) {

	// Read contents from settings file.
	var contents []byte
	if contents, err = ioutil.ReadFile(filename); nil != err {
		return err
	}

	// Parse contents into settings.
	err = yaml.Unmarshal(contents, &Get)
	return
}
