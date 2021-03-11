package refer

import (
	crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

/*
Managed references that in addition to keeping and locating references can also manage their lifecycle:

Auto-creation of missing component using available factories
Auto-linking newly added components
Auto-opening newly added components
Auto-closing removed components
*/
type ManagedReferences struct {
	ReferencesDecorator
	References *crefer.References
	Builder    *BuildReferencesDecorator
	Linker     *LinkReferencesDecorator
	Runner     *RunReferencesDecorator
}

// Creates a new instance of the references
// Parameters:
//   - tuples []interface{}
//   tuples where odd values are component locators (descriptors) and even values are component references
// Returns *ManagedReferences
func NewManagedReferences(tuples []interface{}) *ManagedReferences {
	c := &ManagedReferences{
		ReferencesDecorator: *NewReferencesDecorator(nil, nil),
	}

	c.References = crefer.NewReferences(tuples)
	c.Builder = NewBuildReferencesDecorator(c.References, c)
	c.Linker = NewLinkReferencesDecorator(c.Builder, c)
	c.Runner = NewRunReferencesDecorator(c.Linker, c)

	c.ReferencesDecorator.NextReferences = c.Runner

	return c
}

// Creates a new instance of the references
// Returns *ManagedReferences
func NewEmptyManagedReferences() *ManagedReferences {
	return NewManagedReferences([]interface{}{})
}

// Creates a new ManagedReferences object filled with provided key-value pairs called tuples. Tuples parameters contain a sequence of locator1, component1, locator2, component2, ... pairs.
// Parameters:
//   - tuples ...interface{}
//   the tuples to fill a new ManagedReferences object.
// Returns *ManagedReferences
// a new ManagedReferences object.
func NewManagedReferencesFromTuples(tuples ...interface{}) *ManagedReferences {
	return NewManagedReferences(tuples)
}

// Checks if the component is opened.
// Returns bool
// true if the component has been opened and false otherwise.
func (c *ManagedReferences) IsOpen() bool {
	return c.Linker.IsOpen() && c.Runner.IsOpen()
}

// Opens the component.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
// Returns error
func (c *ManagedReferences) Open(correlationId string) error {
	err := c.Linker.Open(correlationId)
	if err == nil {
		err = c.Runner.Open(correlationId)
	}
	return err
}

// Closes component and frees used resources.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
// Returns error
func (c *ManagedReferences) Close(correlationId string) error {
	err := c.Runner.Close(correlationId)
	if err == nil {
		err = c.Linker.Close(correlationId)
	}
	return err
}
