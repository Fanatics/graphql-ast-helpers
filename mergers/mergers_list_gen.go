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

// SimilarList merges declarations of List that share the same List value.
// This uses the default basic merge strategy.
func SimilarList(curr []*ast.List, more ...*ast.List) ([]*ast.List, error) {
	return Basic.SimilarList(curr, more...)
}

// OneList attempts to merge all members of List into a singe *ast.List.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneList(curr []*ast.List, more ...*ast.List) (*ast.List, error) {
	return Basic.OneList(curr, more...)
}

// SimilarList merges declarations of List that share the same List value.
func (m *Merger) SimilarList(curr []*ast.List, more ...*ast.List) ([]*ast.List, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.List)
	for _, one := range all {
		if one == nil {
			continue
		}
		if key := m.getNodeID(one); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.List
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneList(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneList attempts to merge all members of List into a singe *ast.List.
// If this cannot be done, this method will return an error.
func (m *Merger) OneList(curr []*ast.List, more ...*ast.List) (*ast.List, error) {
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
	var listType []ast.Type
	// range over the parent struct and collect properties
	for _, one := range all {
		listType = append(listType, one.Type)
	}

	var errSet error

	// merge properties

	one := ast.NewList(nil)
	if merged, err := m.OneType(listType); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Type = merged
	}

	return one, errSet

}
