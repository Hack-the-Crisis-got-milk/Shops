package config

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/entities"
	"github.com/Hack-the-Crisis-got-milk/Shops/environment"
	"go.uber.org/config"
)

var (
	// Paths to the config files from the the project's root folder
	baseConfigFile = "./config/base.yaml"
	devConfigFile  = "./config/development.yaml"
	prodConfigFile = "./config/production.yaml"
)

// AppConfig is a struct to store non-private configuration for the project
type AppConfig struct {
	ItemGroups              []entities.ItemGroup `yaml:"item_groups"`
	FeedbackServiceEndpoint string               `yaml:"feedback_endpoint"`
}

// NewAppConfig loads the project config from the config files based on the environment
func NewAppConfig(env *environment.Env) (*AppConfig, error) {
	var configProvider *config.YAML
	var err error
	configFiles := []config.YAMLOption{config.File(baseConfigFile)}
	if env.Get(environment.Environment) == "prod" {
		configFiles = append(configFiles, config.File(prodConfigFile))
	} else if env.Get(environment.Environment) == "dev" {
		configFiles = append(configFiles, config.File(devConfigFile))
	}
	configProvider, err = config.NewYAML(configFiles...)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig

	err = configProvider.Get("").Populate(&cfg)
	return &cfg, nil
}
