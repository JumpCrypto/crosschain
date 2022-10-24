package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func getViper() *viper.Viper {
	// new instance of viper to avoid conflicts with, e.g., cosmos
	v := viper.New()
	v.SetConfigName("crosschain")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("..")
	return v
}

// RequireConfig returns the config - panic if config file is not available
func RequireConfig(section string) map[string]interface{} {
	v := getViper()
	// config is where we store default values
	// panic if not available
	err := v.ReadInConfig()
	if err != nil {
		// fmt.Printf("error reading config file: %w \n", err)
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	// retrieve config
	config := v.GetStringMap(section)
	return config
}

// GetSecret returns a secret, e.g. from env variable. Extend as needed.
func GetSecret(uri string) (string, error) {
	value := uri

	splits := strings.Split(value, ":")
	if len(splits) != 2 {
		return "", errors.New("invalid secret source for: ***")
	}

	path := splits[1]
	switch key := splits[0]; key {
	case "env":
		return strings.TrimSpace(os.Getenv(path)), nil
	case "file":
		if len(path) > 1 && path[0] == '~' {
			path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
		}
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			return "", err
		}
		result, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(result)), nil
	}
	return "", errors.New("invalid secret source for: ***")
}
