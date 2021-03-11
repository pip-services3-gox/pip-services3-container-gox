package examples

import "github.com/pip-services3-gox/pip-services3-container-gox/container"

func NewDummyProcess() *container.ProcessContainer {
	c := container.NewProcessContainer("dummy", "Sample dummy process")
	c.SetConfigPath("./examples/dummy.yaml")
	c.AddFactory(NewDummyFactory())
	return c
}
