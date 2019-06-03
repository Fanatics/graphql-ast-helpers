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

// SimilarVariable merges declarations of Variable that share the same Variable value.
// This uses the default basic merge strategy.
func SimilarVariable(curr []*ast.Variable, more ...*ast.Variable) ([]*ast.Variable, error) {
	return Basic.SimilarVariable(curr, more...)
}

// OneVariable attempts to merge all members of Variable into a singe *ast.Variable.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneVariable(curr []*ast.Variable, more ...*ast.Variable) (*ast.Variable, error) {
	return Basic.OneVariable(curr, more...)
}

// SimilarVariable merges declarations of Variable that share the same Variable value.
func (m *Merger) SimilarVariable(curr []*ast.Variable, more ...*ast.Variable) ([]*ast.Variable, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.Variable)
	for _, one := range all {
		if key := fmt.Sprint(printer.Print(one.Name)); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.Variable
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneVariable(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneVariable attempts to merge all members of Variable into a singe *ast.Variable.
// If this cannot be done, this method will return an error.
func (m *Merger) OneVariable(curr []*ast.Variable, more ...*ast.Variable) (*ast.Variable, error) {
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
	// range over the parent struct and collect properties
	for _, one := range all {
		listName = append(listName, one.Name)
	}

	var errSet error

	// merge properties

	one := ast.NewVariable(nil)
	if merged, err := m.OneName(listName); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Name = merged
	}

	return one, errSet

}
