package test_refer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
	crefer "github.com/pip-services3-gox/pip-services3-container-gox/refer"
)

func TestAutoCreateComponent(t *testing.T) {
	refs := crefer.NewEmptyManagedReferences()

	factory := log.NewDefaultLoggerFactory()
	refs.Put(context.Background(), nil, factory)

	logger, err := refs.GetOneRequired(
		refer.NewDescriptor("*", "logger", "*", "*", "*"),
	)

	assert.Nil(t, err)
	assert.NotNil(t, logger)
}

func TestStringLocator(t *testing.T) {
	refs := crefer.NewEmptyManagedReferences()

	factory := log.NewDefaultLoggerFactory()
	refs.Put(context.Background(), nil, factory)

	logger := refs.GetOneOptional("ABC")

	assert.Nil(t, logger)
}

func TestNilLocator(t *testing.T) {
	refs := crefer.NewEmptyManagedReferences()

	factory := log.NewDefaultLoggerFactory()
	refs.Put(context.Background(), nil, factory)

	logger := refs.GetOneOptional(nil)

	assert.Nil(t, logger)
}
