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

// SimilarVariableDefinition merges declarations of VariableDefinition that share the same VariableDefinition value.
func (m *Merger) SimilarVariableDefinition(curr []*ast.VariableDefinition, more ...*ast.VariableDefinition) ([]*ast.VariableDefinition, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.VariableDefinition)
	for _, one := range all {
		if key := m.getNodeID(one); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.VariableDefinition
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneVariableDefinition(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneVariableDefinition attempts to merge all members of VariableDefinition into a singe *ast.VariableDefinition.
// If this cannot be done, this method will return an error.
func (m *Merger) OneVariableDefinition(curr []*ast.VariableDefinition, more ...*ast.VariableDefinition) (*ast.VariableDefinition, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}
	// prepare property collections
	var listVariable []*ast.Variable
	var listType []ast.Type
	var listDefaultValue []ast.Value
	// range over the parent struct and collect properties
	for _, one := range all {
		listVariable = append(listVariable, one.Variable)
		listType = append(listType, one.Type)
		listDefaultValue = append(listDefaultValue, one.DefaultValue)
	}

	var errSet error

	// merge properties

	one := ast.NewVariableDefinition(nil)
	if merged, err := m.OneVariable(listVariable); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Variable = merged
	}
	if merged, err := m.OneType(listType); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Type = merged
	}
	if merged, err := m.OneValue(listDefaultValue); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.DefaultValue = merged
	}

	return one, errSet

}