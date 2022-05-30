package build

// Creates default container components (loggers, counters, caches, locks, etc.) by their descriptors.
import (
	"github.com/pip-services3-gox/pip-services3-components-gox/auth"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	"github.com/pip-services3-gox/pip-services3-components-gox/cache"
	"github.com/pip-services3-gox/pip-services3-components-gox/config"
	"github.com/pip-services3-gox/pip-services3-components-gox/connect"
	"github.com/pip-services3-gox/pip-services3-components-gox/count"
	"github.com/pip-services3-gox/pip-services3-components-gox/info"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
	"github.com/pip-services3-gox/pip-services3-components-gox/test"
	"github.com/pip-services3-gox/pip-services3-components-gox/trace"
)

// NewDefaultContainerFactory create a new instance of the factory and sets nested factories.
//	Returns: *DefaultContainerFactory
func NewDefaultContainerFactory() *cbuild.CompositeFactory {
	c := cbuild.NewCompositeFactory()

	c.Add(info.NewDefaultInfoFactory())
	c.Add(log.NewDefaultLoggerFactory())
	c.Add(count.NewDefaultCountersFactory())
	c.Add(config.NewDefaultConfigReaderFactory())
	c.Add(cache.NewDefaultCacheFactory())
	c.Add(auth.NewDefaultCredentialStoreFactory())
	c.Add(connect.NewDefaultDiscoveryFactory())
	c.Add(trace.NewDefaultTracerFactory())
	c.Add(log.NewDefaultLoggerFactory())
	c.Add(test.NewDefaultTestFactory())

	return c
}

// NewDefaultContainerFactoryFromFactories create a new instance of the factory and sets nested factories.
//	Parameters:
//		- ctx context.Context
//		- factories ...cbuild.IFactory a list of nested factories
//	Returns: *cbuild.CompositeFactory
func NewDefaultContainerFactoryFromFactories(factories ...cbuild.IFactory) *cbuild.CompositeFactory {
	c := NewDefaultContainerFactory()

	for _, factory := range factories {
		c.Add(factory)
	}

	return c
}
