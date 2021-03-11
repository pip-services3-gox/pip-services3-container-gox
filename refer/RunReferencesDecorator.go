package refer

import (
	crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-commons-go/run"
)

/*

References decorator that automatically opens to newly added components that implement IOpenable interface and
closes removed components that implement ICloseable interface.
*/
type RunReferencesDecorator struct {
	ReferencesDecorator
	opened bool
}

// Creates a new instance of the decorator.
// Parameters:
//   - nextReferences crefer.IReferences
//   the next references or decorator in the chain.
//   - topReferences crefer.IReferences
//   the decorator at the top of the chain.
// Returns *RunReferencesDecorator
func NewRunReferencesDecorator(nextReferences crefer.IReferences,
	topReferences crefer.IReferences) *RunReferencesDecorator {
	return &RunReferencesDecorator{
		ReferencesDecorator: *NewReferencesDecorator(nextReferences, topReferences),
	}
}

// Checks if the component is opened.
// Returns bool
// true if the component has been opened and false otherwise.
func (c *RunReferencesDecorator) IsOpen() bool {
	return c.opened
}

// Opens the component.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
// Returns error
func (c *RunReferencesDecorator) Open(correlationId string) error {
	if !c.opened {
		components := c.GetAll()
		err := run.Opener.Open(correlationId, components)
		c.opened = err == nil
		return err
	}
	return nil
}

// Closes component and frees used resources.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
// Returns error
func (c *RunReferencesDecorator) Close(correlationId string) error {
	if c.opened {
		components := c.GetAll()
		err := run.Closer.Close(correlationId, components)
		c.opened = false
		return err
	}
	return nil
}

// Puts a new reference into this reference map.
// Parameters:
//   - locator interface{}
//   a locator to find the reference by.
//   - component interface{}
//   a component reference to be added.
func (c *RunReferencesDecorator) Put(locator interface{}, component interface{}) {
	c.ReferencesDecorator.Put(locator, component)

	if c.opened {
		run.Opener.OpenOne("", component)
	}
}

// Removes a previously added reference that matches specified locator. If many references match the locator, it removes only the first one. When all references shall be removed, use removeAll method instead.
// see
// removeAll
// Parameters:
//   - locator interface{}
//   a locator to remove reference
// Returns interfce{}
// the removed component reference.
func (c *RunReferencesDecorator) Remove(locator interface{}) interface{} {
	component := c.ReferencesDecorator.Remove(locator)

	if c.opened {
		run.Closer.CloseOne("", component)
	}

	return component
}

// Removes all component references that match the specified locator.
// Parameters:
//   - locator interface{}
//   the locator to remove references by.
// Returns []interface{}
// a list, containing all removed references.
func (c *RunReferencesDecorator) RemoveAll(locator interface{}) []interface{} {
	components := c.NextReferences.RemoveAll(locator)

	if c.opened {
		run.Closer.Close("", components)
	}

	return components
}
