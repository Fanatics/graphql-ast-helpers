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

// SimilarObjectValue merges declarations of ObjectValue that share the same ObjectValue value.
func (m *Merger) SimilarObjectValue(curr []*ast.ObjectValue, more ...*ast.ObjectValue) ([]*ast.ObjectValue, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.ObjectValue)
	for _, one := range all {
		name := m.getValueID(one)
		if name != "" {
			curr, _ := groups[name]
			groups[name] = append(curr, one)
		}
	}

	var out []*ast.ObjectValue
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneObjectValue(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneObjectValue attempts to merge all members of ObjectValue into a singe *ast.ObjectValue.
// If this cannot be done, this method will return an error.
func (m *Merger) OneObjectValue(curr []*ast.ObjectValue, more ...*ast.ObjectValue) (*ast.ObjectValue, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}

	// step 2 - prepare property collections (if any)
  var listFields []*ast.ObjectField

	// step 3 - range over the parent struct and collect properties
	for _, one := range all {
    listFields = append(listFields, one.Fields...)
	}

	// step 4 - prepare output types
	one := ast.NewObjectValue(nil)
	var errSet error

	// step 5 - merge properties
  if merged, err := m.SimilarObjectField(listFields); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Fields = merged
	}

	return one, errSet
}