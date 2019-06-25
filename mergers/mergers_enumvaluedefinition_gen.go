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

// SimilarEnumValueDefinition merges declarations of EnumValueDefinition that share the same EnumValueDefinition value.
// This uses the default basic merge strategy.
func SimilarEnumValueDefinition(curr []*ast.EnumValueDefinition, more ...*ast.EnumValueDefinition) ([]*ast.EnumValueDefinition, error) {
	return Basic.SimilarEnumValueDefinition(curr, more...)
}

// OneEnumValueDefinition attempts to merge all members of EnumValueDefinition into a singe *ast.EnumValueDefinition.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneEnumValueDefinition(curr []*ast.EnumValueDefinition, more ...*ast.EnumValueDefinition) (*ast.EnumValueDefinition, error) {
	return Basic.OneEnumValueDefinition(curr, more...)
}

// SimilarEnumValueDefinition merges declarations of EnumValueDefinition that share the same EnumValueDefinition value.
func (m *Merger) SimilarEnumValueDefinition(curr []*ast.EnumValueDefinition, more ...*ast.EnumValueDefinition) ([]*ast.EnumValueDefinition, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.EnumValueDefinition)
	for _, one := range all {
		if one == nil {
			continue
		}
		if key := fmt.Sprint(printer.Print(one.Name)); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.EnumValueDefinition
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneEnumValueDefinition(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneEnumValueDefinition attempts to merge all members of EnumValueDefinition into a singe *ast.EnumValueDefinition.
// If this cannot be done, this method will return an error.
func (m *Merger) OneEnumValueDefinition(curr []*ast.EnumValueDefinition, more ...*ast.EnumValueDefinition) (*ast.EnumValueDefinition, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	// escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}
	// prepare property collections
	var listName []*ast.Name
	var listDescription []*ast.StringValue
	var listDirectives []*ast.Directive
	// range over the parent struct and collect properties
	for _, one := range all {
		listName = append(listName, one.Name)
		listDescription = append(listDescription, one.Description)
		listDirectives = append(listDirectives, one.Directives...)
	}

	var errSet error

	// merge properties

	one := ast.NewEnumValueDefinition(nil)
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
	if merged, err := m.SimilarDirective(listDirectives); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Directives = merged
	}

	return one, errSet

}
