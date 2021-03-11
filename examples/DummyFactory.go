package examples

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

var ControllerDescriptor = refer.NewDescriptor("pip-services-dummies", "controller", "default", "*", "1.0")

func NewDummyFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ControllerDescriptor, NewDummyController)

	return factory
}
