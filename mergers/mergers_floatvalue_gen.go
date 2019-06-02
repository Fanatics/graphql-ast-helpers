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

// SimilarFloatValue merges declarations of FloatValue that share the same FloatValue value.
func (m *Merger) SimilarFloatValue(curr []*ast.FloatValue, more ...*ast.FloatValue) ([]*ast.FloatValue, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.FloatValue)
	for _, one := range all {
		name := m.getValueID(one)
		if name != "" {
			curr, _ := groups[name]
			groups[name] = append(curr, one)
		}
	}

	var out []*ast.FloatValue
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneFloatValue(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneFloatValue attempts to merge all members of FloatValue into a singe *ast.FloatValue.
// If this cannot be done, this method will return an error.
func (m *Merger) OneFloatValue(curr []*ast.FloatValue, more ...*ast.FloatValue) (*ast.FloatValue, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}

	// step 2 - prepare property collections (if any)

	// step 3 - range over the parent struct and collect properties
	for _, one := range all {
		// 3.a - prevent empty loop from making syntax errors
		_ = one

		// 3.b - accrue properties
	}

	// step 4 - prepare output types
	one := ast.NewFloatValue(nil)
	var errSet error

	// step 5 - merge properties

	return one, errSet
}
