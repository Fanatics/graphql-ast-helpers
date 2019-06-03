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

// SimilarTypeExtensionDefinition merges declarations of TypeExtensionDefinition that share the same TypeExtensionDefinition value.
func (m *Merger) SimilarTypeExtensionDefinition(curr []*ast.TypeExtensionDefinition, more ...*ast.TypeExtensionDefinition) ([]*ast.TypeExtensionDefinition, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.TypeExtensionDefinition)
	for _, one := range all {
		if key := fmt.Sprint(printer.Print(one.Definition.Name)); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.TypeExtensionDefinition
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneTypeExtensionDefinition(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneTypeExtensionDefinition attempts to merge all members of TypeExtensionDefinition into a singe *ast.TypeExtensionDefinition.
// If this cannot be done, this method will return an error.
func (m *Merger) OneTypeExtensionDefinition(curr []*ast.TypeExtensionDefinition, more ...*ast.TypeExtensionDefinition) (*ast.TypeExtensionDefinition, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}
	// prepare property collections
	var listDefinition []*ast.ObjectDefinition
	// range over the parent struct and collect properties
	for _, one := range all {
		listDefinition = append(listDefinition, one.Definition)
	}

	var errSet error

	// merge properties

	one := ast.NewTypeExtensionDefinition(nil)
	if merged, err := m.OneObjectDefinition(listDefinition); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Definition = merged
	}

	return one, errSet

}
