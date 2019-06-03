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

// SimilarInlineFragment merges declarations of InlineFragment that share the same InlineFragment value.
// This uses the default basic merge strategy.
func SimilarInlineFragment(curr []*ast.InlineFragment, more ...*ast.InlineFragment) ([]*ast.InlineFragment, error) {
	return Basic.SimilarInlineFragment(curr, more...)
}

// OneInlineFragment attempts to merge all members of InlineFragment into a singe *ast.InlineFragment.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneInlineFragment(curr []*ast.InlineFragment, more ...*ast.InlineFragment) (*ast.InlineFragment, error) {
	return Basic.OneInlineFragment(curr, more...)
}

// SimilarInlineFragment merges declarations of InlineFragment that share the same InlineFragment value.
func (m *Merger) SimilarInlineFragment(curr []*ast.InlineFragment, more ...*ast.InlineFragment) ([]*ast.InlineFragment, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.InlineFragment)
	for _, one := range all {
		if key := m.getNodeID(one); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.InlineFragment
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneInlineFragment(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneInlineFragment attempts to merge all members of InlineFragment into a singe *ast.InlineFragment.
// If this cannot be done, this method will return an error.
func (m *Merger) OneInlineFragment(curr []*ast.InlineFragment, more ...*ast.InlineFragment) (*ast.InlineFragment, error) {
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
	var listTypeCondition []*ast.Named
	var listDirectives []*ast.Directive
	var listSelectionSet []*ast.SelectionSet
	// range over the parent struct and collect properties
	for _, one := range all {
		listTypeCondition = append(listTypeCondition, one.TypeCondition)
		listDirectives = append(listDirectives, one.Directives...)
		listSelectionSet = append(listSelectionSet, one.SelectionSet)
	}

	var errSet error

	// merge properties

	one := ast.NewInlineFragment(nil)
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
