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

// SimilarIntValue merges declarations of IntValue that share the same IntValue value.
// This uses the default basic merge strategy.
func SimilarIntValue(curr []*ast.IntValue, more ...*ast.IntValue) ([]*ast.IntValue, error) {
	return Basic.SimilarIntValue(curr, more...)
}

// OneIntValue attempts to merge all members of IntValue into a singe *ast.IntValue.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneIntValue(curr []*ast.IntValue, more ...*ast.IntValue) (*ast.IntValue, error) {
	return Basic.OneIntValue(curr, more...)
}

// SimilarIntValue merges declarations of IntValue that share the same IntValue value.
func (m *Merger) SimilarIntValue(curr []*ast.IntValue, more ...*ast.IntValue) ([]*ast.IntValue, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.IntValue)
	for _, one := range all {
		if one == nil {
			continue
		}
		if key := m.getNodeID(one); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.IntValue
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneIntValue(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneIntValue attempts to merge all members of IntValue into a singe *ast.IntValue.
// If this cannot be done, this method will return an error.
func (m *Merger) OneIntValue(curr []*ast.IntValue, more ...*ast.IntValue) (*ast.IntValue, error) {
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

	one := ast.NewIntValue(nil)
	if merged, err := m.Onestring(listValue); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Value = merged
	}

	return one, errSet

}
