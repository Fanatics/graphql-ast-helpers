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

// SimilarFragmentDefinition merges declarations of FragmentDefinition that share the same FragmentDefinition value.
// This uses the default basic merge strategy.
func SimilarFragmentDefinition(curr []*ast.FragmentDefinition, more ...*ast.FragmentDefinition) ([]*ast.FragmentDefinition, error) {
	return Basic.SimilarFragmentDefinition(curr, more...)
}

// OneFragmentDefinition attempts to merge all members of FragmentDefinition into a singe *ast.FragmentDefinition.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneFragmentDefinition(curr []*ast.FragmentDefinition, more ...*ast.FragmentDefinition) (*ast.FragmentDefinition, error) {
	return Basic.OneFragmentDefinition(curr, more...)
}

// SimilarFragmentDefinition merges declarations of FragmentDefinition that share the same FragmentDefinition value.
func (m *Merger) SimilarFragmentDefinition(curr []*ast.FragmentDefinition, more ...*ast.FragmentDefinition) ([]*ast.FragmentDefinition, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.FragmentDefinition)
	for _, one := range all {
		if one == nil {
			continue
		}
		if key := fmt.Sprint(printer.Print(one.Name)); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.FragmentDefinition
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneFragmentDefinition(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneFragmentDefinition attempts to merge all members of FragmentDefinition into a singe *ast.FragmentDefinition.
// If this cannot be done, this method will return an error.
func (m *Merger) OneFragmentDefinition(curr []*ast.FragmentDefinition, more ...*ast.FragmentDefinition) (*ast.FragmentDefinition, error) {
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
	var listOperation []string
	var listName []*ast.Name
	var listVariableDefinitions []*ast.VariableDefinition
	var listTypeCondition []*ast.Named
	var listDirectives []*ast.Directive
	var listSelectionSet []*ast.SelectionSet
	// range over the parent struct and collect properties
	for _, one := range all {
		listOperation = append(listOperation, one.Operation)
		listName = append(listName, one.Name)
		listVariableDefinitions = append(listVariableDefinitions, one.VariableDefinitions...)
		listTypeCondition = append(listTypeCondition, one.TypeCondition)
		listDirectives = append(listDirectives, one.Directives...)
		listSelectionSet = append(listSelectionSet, one.SelectionSet)
	}

	var errSet error

	// merge properties

	one := ast.NewFragmentDefinition(nil)
	if merged, err := m.Onestring(listOperation); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Operation = merged
	}
	if merged, err := m.OneName(listName); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Name = merged
	}
	if merged, err := m.SimilarVariableDefinition(listVariableDefinitions); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.VariableDefinitions = merged
	}
	if merged, err := m.OneNamed(listTypeCondition); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.TypeCondition = merged
	}
	if merged, err := m.SimilarDirective(listDirectives); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Directives = merged
	}
	if merged, err := m.OneSelectionSet(listSelectionSet); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.SelectionSet = merged
	}

	return one, errSet

}
