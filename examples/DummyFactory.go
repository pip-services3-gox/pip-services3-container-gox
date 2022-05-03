package examples

import (
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/build"
)

var ControllerDescriptor = refer.NewDescriptor("pip-services-dummies", "controller", "default", "*", "1.0")

func NewDummyFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ControllerDescriptor, NewDummyController)

	return factory
}
