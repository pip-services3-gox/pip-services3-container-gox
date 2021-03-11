package config

import (
	"path/filepath"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	cconfig "github.com/pip-services3-go/pip-services3-components-go/config"
)

/*
Helper class that reads container configuration from JSON or YAML file.
*/
type TContainerConfigReader struct{}

var ContainerConfigReader = &TContainerConfigReader{}

// Reads container configuration from JSON or YAML file. The type of the file is determined by file extension.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - path string
//  a path to component configuration file.
//  - parameters *config.ConfigParams
//  values to parameters the configuration or null to skip parameterization.
// Returns ContainerConfig, error
// the read container configuration and error
func (c *TContainerConfigReader) ReadFromFile(correlationId string,
	path string, parameters *config.ConfigParams) (ContainerConfig, error) {
	if path == "" {
		return nil, errors.NewConfigError(correlationId, "NO_PATH", "Missing config file path")
	}

	ext := filepath.Ext(path)

	if ext == ".json" {
		return c.ReadFromJsonFile(correlationId, path, parameters)
	}

	if ext == ".yaml" || ext == ".yml" {
		return c.ReadFromYamlFile(correlationId, path, parameters)
	}

	return c.ReadFromJsonFile(correlationId, path, parameters)
}

// Reads container configuration from JSON file.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - path string
//  a path to component configuration file.
//  - parameters *config.ConfigParams
//  values to parameters the configuration or null to skip parameterization.
// Returns ContainerConfig, error
// the read container configuration and error
func (c *TContainerConfigReader) ReadFromJsonFile(correlationId string,
	path string, parameters *config.ConfigParams) (ContainerConfig, error) {
	config, err := cconfig.ReadJsonConfig(correlationId, path, parameters)
	if err != nil {
		return nil, err
	}
	return ReadContainerConfigFromConfig(config)
}

// Reads container configuration from YAML file.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - path string
//  a path to component configuration file.
//  - parameters *config.ConfigParams
//  values to parameters the configuration or null to skip parameterization.
// Returns ContainerConfig, error
// the read container configuration and error
func (c *TContainerConfigReader) ReadFromYamlFile(correlationId string,
	path string, parameters *config.ConfigParams) (ContainerConfig, error) {
	config, err := cconfig.ReadYamlConfig(correlationId, path, parameters)
	if err != nil {
		return nil, err
	}
	return ReadContainerConfigFromConfig(config)
}
