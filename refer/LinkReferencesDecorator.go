package refer

import (
	crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

/*
References decorator that automatically sets references to newly added components that implement IReferenceable
interface and unsets references from removed components that implement IUnreferenceable interface.
*/
type LinkReferencesDecorator struct {
	ReferencesDecorator
	opened bool
}

// Creates a new instance of the decorator.
// Parameters:
//   - nextReferences crefer.IReferences
//   the next references or decorator in the chain.
//   - topReferences crefer.IReferences
//   the decorator at the top of the chain.
// Returns *LinkReferencesDecorator
func NewLinkReferencesDecorator(nextReferences crefer.IReferences,
	topReferences crefer.IReferences) *LinkReferencesDecorator {
	return &LinkReferencesDecorator{
		ReferencesDecorator: *NewReferencesDecorator(nextReferences, topReferences),
	}
}

// Checks if the component is opened.
// Returns bool
// true if the component has been opened and false otherwise.
func (c *LinkReferencesDecorator) IsOpen() bool {
	return c.opened
}

// Opens the component.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
// Returns error
func (c *LinkReferencesDecorator) Open(correlationId string) error {
	if !c.opened {
		c.opened = true
		components := c.GetAll()
		crefer.Referencer.SetReferences(c.ReferencesDecorator.TopReferences, components)
	}
	return nil
}

// Closes component and frees used resources.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
// Returns error
func (c *LinkReferencesDecorator) Close(correlationId string) error {
	if c.opened {
		c.opened = false
		components := c.GetAll()
		crefer.Referencer.UnsetReferences(components)
	}
	return nil
}

// Puts a new reference into this reference map.
// Parameters:
//   - locator intrface{}
//   a locator to find the reference by.
//   - component interface{}
//   a component reference to be added.
func (c *LinkReferencesDecorator) Put(locator interface{}, component interface{}) {
	c.ReferencesDecorator.Put(locator, component)

	if c.opened {
		crefer.Referencer.SetReferencesForOne(c.ReferencesDecorator.TopReferences, component)
	}
}

// Removes a previously added reference that matches specified locator. If many references match the locator, it removes only the first one.
// When all references shall be removed, use removeAll method instead.
// see
// removeAll
// Parameters:
//   - locator interface
//   a locator to remove reference
// Returns interface{}
// the removed component reference.
func (c *LinkReferencesDecorator) Remove(locator interface{}) interface{} {
	component := c.ReferencesDecorator.Remove(locator)

	if c.opened {
		crefer.Referencer.UnsetReferencesForOne(component)
	}

	return component
}

// Removes all component references that match the specified locator.
// Parameters:
//   - locator interface{}
//   the locator to remove references by.
// Returns []interface{}
// a list, containing all removed references.
func (c *LinkReferencesDecorator) RemoveAll(locator interface{}) []interface{} {
	components := c.NextReferences.RemoveAll(locator)

	if c.opened {
		crefer.Referencer.UnsetReferences(components)
	}

	return components
}
