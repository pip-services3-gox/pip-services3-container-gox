package refer

import crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"

/*
Chainable decorator for IReferences that allows to inject additional capabilities such as
automatic component creation, automatic registration and opening.
*/

type ReferencesDecorator struct {
	NextReferences crefer.IReferences
	TopReferences  crefer.IReferences
}

// Creates a new instance of the decorator.
// Parameters:
//   - nextReferences crefer.IReferences
//   the next references or decorator in the chain.
//   - topReferences crefer.IReferences
//   the decorator at the top of the chain.
// Returns *ReferencesDecorator
func NewReferencesDecorator(nextReferences crefer.IReferences,
	topReferences crefer.IReferences) *ReferencesDecorator {
	c := &ReferencesDecorator{
		NextReferences: nextReferences,
		TopReferences:  topReferences,
	}

	if c.NextReferences == nil {
		c.NextReferences = topReferences
	}
	if c.TopReferences == nil {
		c.TopReferences = nextReferences
	}

	return c
}

// Puts a new reference into this reference map.
// Parameters:
//   - locator interface{}
//   a locator to find the reference by.
//   - component interface{}
//   a component reference to be added.
func (c *ReferencesDecorator) Put(locator interface{}, component interface{}) {
	c.NextReferences.Put(locator, component)
}

// Removes a previously added reference that matches specified locator. If many references match the locator, it removes only the first one. When all references shall be removed, use removeAll method instead.
// see
// RemoveAll
// Parameters:
//   - locator  interface{}
//   a locator to remove reference
// Returns interface{}
// the removed component reference.
func (c *ReferencesDecorator) Remove(locator interface{}) interface{} {
	return c.NextReferences.Remove(locator)
}

// Removes all component references that match the specified locator.
// Parameters:
//   - locator interface{}
//   the locator to remove references by.
// Returns []interface{}
// a list, containing all removed references.
func (c *ReferencesDecorator) RemoveAll(locator interface{}) []interface{} {
	return c.NextReferences.RemoveAll(locator)
}

// Gets locators for all registered component references in this reference map.
// Returns []interface{}
// a list with component locators.
func (c *ReferencesDecorator) GetAllLocators() []interface{} {
	return c.NextReferences.GetAllLocators()
}

// Gets all component references registered in this reference map.
// Returns []interface{}
// a list with component references.
func (c *ReferencesDecorator) GetAll() []interface{} {
	return c.NextReferences.GetAll()
}

// Gets an optional component reference that matches specified locator.
// Parameters:
//   - locator interface{}
//   the locator to find references by.
// Returns interface{}
// a matching component reference or null if nothing was found.
func (c *ReferencesDecorator) GetOneOptional(locator interface{}) interface{} {
	var component interface{}

	defer func() {
		recover()
	}()

	components, err := c.Find(locator, false)
	if err == nil && len(components) > 0 {
		component = components[0]
	}

	return component
}

// Gets a required component reference that matches specified locator.
// Parameters:
//   - locator interface{}
//   the locator to find a reference by.
// Returns interface{}, error
// a matching component reference, a ReferenceError when no references found.
func (c *ReferencesDecorator) GetOneRequired(locator interface{}) (interface{}, error) {
	components, err := c.Find(locator, true)
	if err != nil || len(components) == 0 {
		return nil, err
	}
	return components[0], nil
}

// Gets all component references that match specified locator.
// Parameters:
//   - locator interface{}
//   the locator to find references by.
// Returns []interface{}
// a list with matching component references or empty list if nothing was found.
func (c *ReferencesDecorator) GetOptional(locator interface{}) []interface{} {
	components := []interface{}{}

	defer func() {
		recover()
	}()

	components, _ = c.Find(locator, false)

	return components
}

// Gets all component references that match specified locator. At least one component reference must be present. If it doesn't the method throws an error.
// Parameters:
//   - locator interface{}
//   the locator to find references by.
// Returns []interface{}
// a list with matching component references and error a ReferenceError when no references found.
func (c *ReferencesDecorator) GetRequired(locator interface{}) ([]interface{}, error) {
	return c.Find(locator, true)
}

// Gets all component references that match specified locator.
// Parameters:
//   - locator interface{}
//   the locator to find a reference by.
//   - required bool
//   forces to raise an exception if no reference is found.
// Returns []interface{}, error
// a list with matching component references and a ReferenceError when required is set to true but no references found
func (c *ReferencesDecorator) Find(locator interface{}, required bool) ([]interface{}, error) {
	return c.NextReferences.Find(locator, required)
}
