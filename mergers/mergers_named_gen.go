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

// SimilarNamed merges declarations of Named that share the same Named value.
func (m *Merger) SimilarNamed(curr []*ast.Named, more ...*ast.Named) ([]*ast.Named, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.Named)
	for _, one := range all {
		name := fmt.Sprint(printer.Print(one))
		if name != "" {
			curr, _ := groups[name]
			groups[name] = append(curr, one)
		}
	}

	var out []*ast.Named
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneNamed(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneNamed attempts to merge all members of Named into a singe *ast.Named.
// If this cannot be done, this method will return an error.
func (m *Merger) OneNamed(curr []*ast.Named, more ...*ast.Named) (*ast.Named, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}

	// step 2 - prepare property collections (if any)
  var listName []*ast.Name

	// step 3 - range over the parent struct and collect properties
	for _, one := range all {
    listName = append(listName, one.Name)
	}

	// step 4 - prepare output types
	one := ast.NewNamed(nil)
	var errSet error

	// step 5 - merge properties
  if merged, err := m.OneName(listName); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Name = merged
	}

	return one, errSet
}