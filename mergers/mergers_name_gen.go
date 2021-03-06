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

// SimilarName merges declarations of Name that share the same Name value.
// This uses the default basic merge strategy.
func SimilarName(curr []*ast.Name, more ...*ast.Name) ([]*ast.Name, error) {
	return Basic.SimilarName(curr, more...)
}

// OneName attempts to merge all members of Name into a singe *ast.Name.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneName(curr []*ast.Name, more ...*ast.Name) (*ast.Name, error) {
	return Basic.OneName(curr, more...)
}

// SimilarName merges declarations of Name that share the same Name value.
func (m *Merger) SimilarName(curr []*ast.Name, more ...*ast.Name) ([]*ast.Name, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.Name)
	for _, one := range all {
		if one == nil {
			continue
		}
		if key := fmt.Sprint(printer.Print(one)); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.Name
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneName(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneName attempts to merge all members of Name into a singe *ast.Name.
// If this cannot be done, this method will return an error.
func (m *Merger) OneName(curr []*ast.Name, more ...*ast.Name) (*ast.Name, error) {
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
	var listValue []string
	// range over the parent struct and collect properties
	for _, one := range all {
		listValue = append(listValue, one.Value)
	}

	var errSet error

	// merge properties

	one := ast.NewName(nil)
	if merged, err := m.Onestring(listValue); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Value = merged
	}

	return one, errSet

}
