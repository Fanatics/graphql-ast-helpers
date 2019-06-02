// Code generated by go generate; DO NOT EDIT.
package mergers

import (
	"fmt"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
	"github.com/richardwilkes/toolbox/errs"
)

var _ = fmt.Sprint
var _ = printer.Print

// SimilarObjectDefinition merges declarations of ObjectDefinition that share the same ObjectDefinition value.
func (m *Merger) SimilarObjectDefinition(curr []*ast.ObjectDefinition, more ...*ast.ObjectDefinition) ([]*ast.ObjectDefinition, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.ObjectDefinition)
	for _, one := range all {
		name := fmt.Sprint(printer.Print(one.Name))
		if name != "" {
			curr, _ := groups[name]
			groups[name] = append(curr, one)
		}
	}

	var out []*ast.ObjectDefinition
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneObjectDefinition(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneObjectDefinition attempts to merge all members of ObjectDefinition into a singe *ast.ObjectDefinition.
// If this cannot be done, this method will return an error.
func (m *Merger) OneObjectDefinition(curr []*ast.ObjectDefinition, more ...*ast.ObjectDefinition) (*ast.ObjectDefinition, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}

	// step 2 - prepare property collections (if any)
  var listName []*ast.Name
  var listDescription []*ast.StringValue
  var listInterfaces []*ast.Named
  var listDirectives []*ast.Directive
  var listFields []*ast.FieldDefinition

	// step 3 - range over the parent struct and collect properties
	for _, one := range all {
    listName = append(listName, one.Name)
    listDescription = append(listDescription, one.Description)
    listInterfaces = append(listInterfaces, one.Interfaces...)
    listDirectives = append(listDirectives, one.Directives...)
    listFields = append(listFields, one.Fields...)
	}

	// step 4 - prepare output types
	one := ast.NewObjectDefinition(nil)
	var errSet error

	// step 5 - merge properties
  if merged, err := m.OneName(listName); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Name = merged
	}
  if merged, err := m.OneStringValue(listDescription); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Description = merged
	}
  if merged, err := m.SimilarNamed(listInterfaces); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Interfaces = merged
	}
  if merged, err := m.SimilarDirective(listDirectives); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Directives = merged
	}
  if merged, err := m.SimilarFieldDefinition(listFields); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Fields = merged
	}

	return one, errSet
}