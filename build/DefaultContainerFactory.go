package build

/*
Creates default container components (loggers, counters, caches, locks, etc.) by their descriptors.
*/
import (
	"github.com/pip-services3-go/pip-services3-components-go/auth"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-components-go/cache"
	"github.com/pip-services3-go/pip-services3-components-go/config"
	"github.com/pip-services3-go/pip-services3-components-go/connect"
	"github.com/pip-services3-go/pip-services3-components-go/count"
	"github.com/pip-services3-go/pip-services3-components-go/info"
	"github.com/pip-services3-go/pip-services3-components-go/log"
	"github.com/pip-services3-go/pip-services3-components-go/test"
)

// Create a new instance of the factory and sets nested factories.
// Returns *DefaultContainerFactory
func NewDefaultContainerFactory() *cbuild.CompositeFactory {
	c := cbuild.NewCompositeFactory()

	c.Add(info.NewDefaultInfoFactory())
	c.Add(log.NewDefaultLoggerFactory())
	c.Add(count.NewDefaultCountersFactory())
	c.Add(config.NewDefaultConfigReaderFactory())
	c.Add(cache.NewDefaultCacheFactory())
	c.Add(auth.NewDefaultCredentialStoreFactory())
	c.Add(connect.NewDefaultDiscoveryFactory())
	c.Add(log.NewDefaultLoggerFactory())
	c.Add(test.NewDefaultTestFactory())

	return c
}

// Create a new instance of the factory and sets nested factories.
// Parameters:
//  - factories ...cbuild.IFactory
//  a list of nested factories
// Returns *cbuild.CompositeFactory
func NewDefaultContainerFactoryFromFactories(factories ...cbuild.IFactory) *cbuild.CompositeFactory {
	c := NewDefaultContainerFactory()

	for _, factory := range factories {
		c.Add(factory)
	}

	return c
}
