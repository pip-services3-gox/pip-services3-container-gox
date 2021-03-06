package container

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/log"
)

/*
Inversion of control (IoC) container that runs as a system process. It processes command line arguments and handles unhandled exceptions and Ctrl-C signal to gracefully shutdown the container.

Command line arguments
  --config / -c path to JSON or YAML file with container configuration (default: "./config/config.yml")
  --param / --params / -p value(s) to parameterize the container configuration
  --help / -h prints the container usage help
see
Container

Example
  container = NewEmptyProcessContainer();
  container.Container.AddFactory(NewMyComponentFactory());

  container.Run(process.args);
*/
type ProcessContainer struct {
	Container
	configPath string
}

// Creates a new empty instance of the container.
// Returns ProcessContainer
func NewEmptyProcessContainer() *ProcessContainer {
	c := &ProcessContainer{
		Container:  *NewEmptyContainer(),
		configPath: "./config/config.yml",
	}
	c.SetLogger(log.NewConsoleLogger())
	return c
}

// Creates a new instance of the container.
// Parameters:
//   - name string
//   a container name (accessible via ContextInfo)
//   - description string
//   a container description (accessible via ContextInfo)
// Returns ProcessContainer
func NewProcessContainer(name string, description string) *ProcessContainer {
	c := &ProcessContainer{
		Container:  *NewContainer(name, description),
		configPath: "./config/config.yml",
	}
	c.SetLogger(log.NewConsoleLogger())
	return c
}

// Creates a new instance of the container inherit from reference.
// Parameters:
//   - name string
//   a container name (accessible via ContextInfo)
//   - description string
//   a container description (accessible via ContextInfo)
//   - referenceable crefer.IReferenceable
//   - referenceble object for inherit
// Returns *Container
func InheritProcessContainer(name string, description string,
	referenceable crefer.IReferenceable) *ProcessContainer {
	c := &ProcessContainer{
		Container:  *InheritContainer(name, description, referenceable),
		configPath: "./config/config.yml",
	}
	c.SetLogger(log.NewConsoleLogger())
	return c
}

// Set path for configuration file
func (c *ProcessContainer) SetConfigPath(configPath string) {
	c.configPath = configPath
}

func (c *ProcessContainer) getConfigPath(args []string) string {
	for index, arg := range args {
		nextArg := ""
		if index < len(args)-1 {
			nextArg = args[index+1]
			if strings.HasPrefix(nextArg, "-") {
				nextArg = ""
			}
		}

		if arg == "--config" || arg == "-c" {
			return nextArg
		}
	}

	return c.configPath
}

func (c *ProcessContainer) getParameters(args []string) *cconfig.ConfigParams {
	line := ""

	for index := 0; index < len(args); index++ {
		arg := args[index]
		nextArg := ""
		if index < len(args)-1 {
			nextArg = args[index+1]
			if strings.HasPrefix(nextArg, "-") {
				nextArg = ""
			}
		}

		if nextArg != "" {
			if arg == "--param" || arg == "--params" || arg == "-p" {
				if line != "" {
					line = line + ";"
				}
				line = line + nextArg
				index++
			}
		}
	}

	parameters := cconfig.NewConfigParamsFromString(line)

	for _, e := range os.Environ() {
		env := strings.Split(e, "=")
		parameters.SetAsObject(env[0], env[1])
	}

	return parameters
}

func (c *ProcessContainer) showHelp(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func (c *ProcessContainer) printHelp() {
	fmt.Println("Pip.Services process container - http://www.github.com/pip-services/pip-services")
	fmt.Println("run [-h] [-c <config file>] [-p <param>=<value>]*")
}

func (c *ProcessContainer) captureErrors(correlationId string) {
	if r := recover(); r != nil {
		err, _ := r.(error)
		c.Logger().Fatal(correlationId, err, "Process is terminated")
		os.Exit(1)
	}
}

func (c *ProcessContainer) captureExit(correlationId string) {
	c.Logger().Info(correlationId, "Press Control-C to stop the microservice...")

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		select {
		case <-ch:
			c.Close(correlationId)
			c.Logger().Info(correlationId, "Goodbye!")
			os.Exit(0)
		}
	}()
}

// Runs the container by instantiating and running components inside the container.
// It reads the container configuration, creates, configures, references and opens components. On process exit it closes, unreferences and destroys components to gracefully shutdown.
// Parameters:
//   - args []string
//   command line arguments
func (c *ProcessContainer) Run(args []string) {
	if c.showHelp(args) {
		c.printHelp()
		os.Exit(0)
		return
	}

	correlationId := c.Info().Name
	path := c.getConfigPath(args)
	parameters := c.getParameters(args)

	err := c.ReadConfigFromFile(correlationId, path, parameters)
	if err != nil {
		c.Logger().Fatal(correlationId, err, "Process is terminated")
		os.Exit(1)
		return
	}

	defer c.captureErrors(correlationId)
	c.captureExit(correlationId)

	err = c.Open(correlationId)
	if err != nil {
		c.Logger().Fatal(correlationId, err, "Process is terminated")
		os.Exit(1)
		return
	}

	for {
		time.Sleep(100)
	}
}
