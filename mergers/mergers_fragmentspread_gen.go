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

// SimilarFragmentSpread merges declarations of FragmentSpread that share the same FragmentSpread value.
func (m *Merger) SimilarFragmentSpread(curr []*ast.FragmentSpread, more ...*ast.FragmentSpread) ([]*ast.FragmentSpread, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.FragmentSpread)
	for _, one := range all {
		if key := fmt.Sprint(printer.Print(one.Name)); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.FragmentSpread
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneFragmentSpread(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneFragmentSpread attempts to merge all members of FragmentSpread into a singe *ast.FragmentSpread.
// If this cannot be done, this method will return an error.
func (m *Merger) OneFragmentSpread(curr []*ast.FragmentSpread, more ...*ast.FragmentSpread) (*ast.FragmentSpread, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}
	// prepare property collections
	var listName []*ast.Name
	var listDirectives []*ast.Directive
	// range over the parent struct and collect properties
	for _, one := range all {
		listName = append(listName, one.Name)
		listDirectives = append(listDirectives, one.Directives...)
	}

	var errSet error

	// merge properties

	one := ast.NewFragmentSpread(nil)
	if merged, err := m.OneName(listName); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Name = merged
	}
	if merged, err := m.SimilarDirective(listDirectives); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Directives = merged
	}

	return one, errSet

}