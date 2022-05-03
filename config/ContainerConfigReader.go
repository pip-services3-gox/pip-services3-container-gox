package config

import (
	"context"
	"path/filepath"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	cconfig "github.com/pip-services3-gox/pip-services3-components-gox/config"
)

// ContainerConfigReader Helper class that reads container configuration from JSON or YAML file.
var ContainerConfigReader = &_TContainerConfigReader{}

type _TContainerConfigReader struct{}

// ReadFromFile reads container configuration from JSON or YAML file.
// The type of the file is determined by file extension.
//	Parameters:
//		- ctx context.Context.
//		- correlationId string transaction id to trace execution through call chain.
//		- path string a path to component configuration file.
//		- parameters *config.ConfigParams values to parameters the configuration or null to skip parameterization.
//	Returns: ContainerConfig, error the read container configuration and error
func (c *_TContainerConfigReader) ReadFromFile(ctx context.Context, correlationId string,
	path string, parameters *config.ConfigParams) (ContainerConfig, error) {
	if path == "" {
		return nil, errors.NewConfigError(correlationId, "NO_PATH", "Missing config file path")
	}

	ext := filepath.Ext(path)

	if ext == ".json" {
		return c.ReadFromJsonFile(ctx, correlationId, path, parameters)
	}

	if ext == ".yaml" || ext == ".yml" {
		return c.ReadFromYamlFile(ctx, correlationId, path, parameters)
	}

	return c.ReadFromJsonFile(ctx, correlationId, path, parameters)
}

// ReadFromJsonFile reads container configuration from JSON file.
//	Parameters:
//		- ctx context.Context.
//		- correlationId string transaction id to trace execution through call chain.
//		- path string a path to component configuration file.
//		- parameters *config.ConfigParams values to parameters the configuration or null to skip parameterization.
//	Returns: ContainerConfig, error the read container configuration and error
func (c *_TContainerConfigReader) ReadFromJsonFile(ctx context.Context, correlationId string,
	path string, parameters *config.ConfigParams) (ContainerConfig, error) {

	config, err := cconfig.ReadJsonConfig(ctx, correlationId, path, parameters)
	if err != nil {
		return nil, err
	}
	return ReadContainerConfigFromConfig(config)
}

// ReadFromYamlFile reads container configuration from YAML file.
//	Parameters:
//		- ctx context.Context.
//		- correlationId string transaction id to trace execution through call chain.
//		- path string a path to component configuration file.
//		- parameters *config.ConfigParams values to parameters the configuration or null to skip parameterization.
//	Returns: ContainerConfig, error the read container configuration and error
func (c *_TContainerConfigReader) ReadFromYamlFile(ctx context.Context, correlationId string,
	path string, parameters *config.ConfigParams) (ContainerConfig, error) {

	config, err := cconfig.ReadYamlConfig(ctx, correlationId, path, parameters)
	if err != nil {
		return nil, err
	}
	return ReadContainerConfigFromConfig(config)
}
