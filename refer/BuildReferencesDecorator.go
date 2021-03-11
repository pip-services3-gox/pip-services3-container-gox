package refer

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
References decorator that automatically creates missing components using available
component factories upon component retrival.
*/
type BuildReferencesDecorator struct {
	ReferencesDecorator
}

// Creates a new instance of the decorator.
// Parameters:
//   - nextReferences crefer.IReferences
//   the next references or decorator in the chain.
//   - topReferences IReferences
//   the decorator at the top of the chain.
// Returns *BuildReferencesDecorator
func NewBuildReferencesDecorator(nextReferences crefer.IReferences,
	topReferences crefer.IReferences) *BuildReferencesDecorator {
	return &BuildReferencesDecorator{
		ReferencesDecorator: *NewReferencesDecorator(nextReferences, topReferences),
	}
}

// Finds a factory capable creating component by given descriptor from the components registered in the
// references.
// Parameters:
//   - locator interface{}
//   a locator of component to be created.
// Returns build.IFactory
// found factory or nil if factory was not found.
func (c *BuildReferencesDecorator) FindFactory(locator interface{}) build.IFactory {
	components := c.GetAll()

	for _, component := range components {
		factory, ok := component.(build.IFactory)
		if ok && factory.CanCreate(locator) != nil {
			return factory
		}
	}

	return nil
}

// Creates a component identified by given locator.
// throws
// a CreateEerror if the factory is not able to create the component.
// see
// findFactory
// Parameters:
//   - locator interface{}
//   a locator to identify component to be created.
//   - factory build.IFactory
//   a factory that shall create the component.
// Returns interface{}
// the created component.
func (c *BuildReferencesDecorator) Create(locator interface{},
	factory build.IFactory) interface{} {

	if factory == nil {
		return nil
	}

	var result interface{}

	defer func() {
		recover()
	}()

	result, _ = factory.Create(locator)

	return result
}

// Clarifies a component locator by merging two descriptors into one to replace missing fields.
// That allows to get a more complete descriptor that includes all possible fields.
// Parameters:
//   - locator intrface{}
//   a component locator to clarify.
//   - factory build.IFactory
//   a factory that shall create the component.
// Returns interface{}
// clarified component descriptor (locator)
func (c *BuildReferencesDecorator) ClarifyLocator(locator interface{},
	factory build.IFactory) interface{} {

	if factory == nil {
		return nil
	}

	descriptor, ok := locator.(*refer.Descriptor)
	if !ok {
		return locator
	}

	anotherLocator := factory.CanCreate(locator)
	anotherDescriptor, ok1 := anotherLocator.(*refer.Descriptor)
	if !ok1 {
		return locator
	}

	group := descriptor.Group()
	if group == "" {
		group = anotherDescriptor.Group()
	}
	typ := descriptor.Type()
	if typ == "" {
		typ = anotherDescriptor.Type()
	}
	kind := descriptor.Kind()
	if kind == "" {
		kind = anotherDescriptor.Kind()
	}
	name := descriptor.Name()
	if name == "" {
		name = anotherDescriptor.Name()
	}
	version := descriptor.Version()
	if version == "" {
		version = anotherDescriptor.Version()
	}

	return refer.NewDescriptor(group, typ, kind, name, version)
}

// Gets an optional component reference that matches specified locator.
// Parameters:
//   - locator interface{}
//   the locator to find references by.
// Returns interface{}
// a matching component reference or nil if nothing was found.
func (c *BuildReferencesDecorator) GetOneOptional(locator interface{}) interface{} {
	components, err := c.Find(locator, false)
	if err != nil || len(components) == 0 {
		return nil
	}
	return components[0]
}

// Gets a required component reference that matches specified locator.
// throws
// a ReferenceException when no references found.
// Parameters:
//   - locator interface{}
//   the locator to find a reference by.
// Returns interface{}, error
// a matching component reference and error.
func (c *BuildReferencesDecorator) GetOneRequired(locator interface{}) (interface{}, error) {
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
func (c *BuildReferencesDecorator) GetOptional(locator interface{}) []interface{} {
	components, _ := c.Find(locator, false)
	return components
}

// Gets all component references that match specified locator. At least one component reference must be present.
// If it doesn't the method throws an error.
// throws
// a ReferenceException when no references found.
// Parameters:
//  - locator interface{}
//  the locator to find references by.
// Returns []interface{}, erorr
// a list with matching component references and error.
func (c *BuildReferencesDecorator) GetRequired(locator interface{}) ([]interface{}, error) {
	return c.Find(locator, true)
}

// Gets all component references that match specified locator.
// throws
// a ReferenceError when required is set to true but no references found.
// Parameters:
//   - locator interface
//   the locator to find a reference by.
//   - required bool
//   forces to raise an exception if no reference is found.
// Returns []interface, error
// a list with matching component references and error.
func (c *BuildReferencesDecorator) Find(locator interface{}, required bool) ([]interface{}, error) {
	components, _ := c.ReferencesDecorator.Find(locator, required)

	if required && len(components) == 0 {
		factory := c.FindFactory(locator)
		component := c.Create(locator, factory)
		if component != nil {
			locator = c.ClarifyLocator(locator, factory)
			c.ReferencesDecorator.TopReferences.Put(locator, component)
			components = append(components, component)
		}
	}

	if required && len(components) == 0 {
		err := refer.NewReferenceError("", locator)
		return nil, err
	}

	return components, nil
}
