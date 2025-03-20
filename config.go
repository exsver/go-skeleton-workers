package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	MaxProcs int `json:"maxProcs" validate:"gte=0,lte=1024"`
	Workers  int `json:"workers" validate:"gte=1,lte=1024"`
}

func (c *Config) String() string {
	return fmt.Sprintf("MaxProcs: %d, Workers: %d",
		c.MaxProcs,
		c.Workers,
	)
}

func getConfiguration(flags *Flags) (*Config, error) {
	config, err := readConfigFile(*flags.ConfigFilePath)
	if err != nil {
		return nil, err
	}

	// set defaults
	if config.MaxProcs == 0 {
		config.MaxProcs = runtime.NumCPU()
	}

	err = config.Validate()
	if err != nil {
		return nil, err
	}

	Log.Info.Printf("Configuration: %s", config.String())

	return config, nil
}

func readConfigFile(filePath string) (*Config, error) {
	Log.Debug.Printf("Reading config file '%s'", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var out Config

	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Config) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
