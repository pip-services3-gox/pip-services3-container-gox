package refer

import (
	"context"
	crefer "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
)

// RunReferencesDecorator References decorator that automatically opens
// to newly added components that implement IOpenable interface and
// closes removed components that implement ICloseable interface.
type RunReferencesDecorator struct {
	ReferencesDecorator
	opened bool
}

// NewRunReferencesDecorator creates a new instance of the decorator.
//	Parameters:
//		- nextReferences crefer.IReferences the next references or decorator in the chain.
//		- topReferences crefer.IReferences the decorator at the top of the chain.
//	Returns: *RunReferencesDecorator
func NewRunReferencesDecorator(nextReferences crefer.IReferences,
	topReferences crefer.IReferences) *RunReferencesDecorator {

	return &RunReferencesDecorator{
		ReferencesDecorator: *NewReferencesDecorator(nextReferences, topReferences),
	}
}

// IsOpen checks if the component is opened.
//	Returns: bool true if the component has been opened and false otherwise.
func (c *RunReferencesDecorator) IsOpen() bool {
	return c.opened
}

// Open the component.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//	Returns: error
func (c *RunReferencesDecorator) Open(ctx context.Context, correlationId string) error {
	if !c.opened {
		components := c.GetAll()
		err := run.Opener.Open(ctx, correlationId, components)
		c.opened = err == nil
		return err
	}
	return nil
}

// Close component and frees used resources.
//	Parameters:
//		- ctx context.Context
//		- correlationId string transaction id to trace execution through call chain.
//	Returns: error
func (c *RunReferencesDecorator) Close(ctx context.Context, correlationId string) error {
	if c.opened {
		components := c.GetAll()
		err := run.Closer.Close(ctx, correlationId, components)
		c.opened = false
		return err
	}
	return nil
}

// Put a new reference into this reference map.
//	Parameters:
//		- locator any a locator to find the reference by.
//		- component any a component reference to be added.
func (c *RunReferencesDecorator) Put(locator any, component any) {
	c.ReferencesDecorator.Put(locator, component)

	if c.opened {
		run.Opener.OpenOne(context.Background(), "", component)
	}
}

// Remove a previously added reference that matches specified locator.
// If many references match the locator, it removes only the first one.
// When all references shall be removed, use removeAll method instead.
//	see RemoveAll
//	Parameters: locator any a locator to remove reference
//	Returns: any the removed component reference.
func (c *RunReferencesDecorator) Remove(locator any) any {
	component := c.ReferencesDecorator.Remove(locator)

	if c.opened {
		run.Closer.CloseOne(context.Background(), "", component)
	}

	return component
}

// RemoveAll all component references that match the specified locator.
//	Parameters: locator any the locator to remove references by.
//	Returns: []any a list, containing all removed references.
func (c *RunReferencesDecorator) RemoveAll(locator any) []any {
	components := c.NextReferences.RemoveAll(locator)

	if c.opened {
		run.Closer.Close(context.Background(), "", components)
	}

	return components
}
