package design

import (
	"fmt"

	"goa.design/goa.v2/eval"
)

type (
	// ServiceExpr describes a set of related methods.
	ServiceExpr struct {
		// DSLFunc contains the DSL used to initialize the expression.
		eval.DSLFunc
		// Name of method group.
		Name string
		// Description of method group for consumption by humans.
		Description string
		// Docs points to external documentation
		Docs *DocsExpr
		// Servers list the API hosts
		Servers []*ServerExpr
		// Methods is the list of service methods.
		Methods []*MethodExpr
		// Errors list the errors common to all the service methods.
		Errors []*ErrorExpr
		// Metadata is a set of key/value pairs with semantic that is
		// specific to each generator.
		Metadata MetadataExpr
	}

	// ErrorExpr defines an error response. It consists of a named
	// attribute.
	ErrorExpr struct {
		// AttributeExpr is the underlying attribute.
		*AttributeExpr
		// Name is the unique name of the error.
		Name string
	}
)

// EvalName returns the generic expression name used in error messages.
func (s *ServiceExpr) EvalName() string {
	if s.Name == "" {
		return "unnamed service"
	}
	return fmt.Sprintf("service %#v", s.Name)
}

// Error returns the error with the given name if any.
func (s *ServiceExpr) Error(name string) *ErrorExpr {
	for _, erro := range s.Errors {
		if erro.Name == name {
			return erro
		}
	}
	return Root.Error(name)
}

// Hash returns a unique hash value for s.
func (s *ServiceExpr) Hash() string {
	return "_service_+" + s.Name
}

// Finalize finalizes all the service methods.
func (s *ServiceExpr) Finalize() {
	for _, ep := range s.Methods {
		ep.Finalize()
	}
}