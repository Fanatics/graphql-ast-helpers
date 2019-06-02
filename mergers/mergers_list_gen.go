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
func (m *Merger) SimilarList(curr []*ast.List, more ...*ast.List) ([]*ast.List, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.List)
	for _, one := range all {
		name := fmt.Sprint(printer.Print(one))
		if name != "" {
			curr, _ := groups[name]
			groups[name] = append(curr, one)
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
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}

	// step 2 - prepare property collections (if any)
  var listType []ast.Type

	// step 3 - range over the parent struct and collect properties
	for _, one := range all {
    listType = append(listType, one.Type)
	}

	// step 4 - prepare output types
	one := ast.NewList(nil)
	var errSet error

	// step 5 - merge properties
  if merged, err := m.OneType(listType); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Type = merged
	}

	return one, errSet
}
