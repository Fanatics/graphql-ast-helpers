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

// SimilarDirective merges declarations of Directive that share the same Directive value.
// This uses the default basic merge strategy.
func SimilarDirective(curr []*ast.Directive, more ...*ast.Directive) ([]*ast.Directive, error) {
	return Basic.SimilarDirective(curr, more...)
}

// OneDirective attempts to merge all members of Directive into a singe *ast.Directive.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneDirective(curr []*ast.Directive, more ...*ast.Directive) (*ast.Directive, error) {
	return Basic.OneDirective(curr, more...)
}

// SimilarDirective merges declarations of Directive that share the same Directive value.
func (m *Merger) SimilarDirective(curr []*ast.Directive, more ...*ast.Directive) ([]*ast.Directive, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.Directive)
	for _, one := range all {
		if one == nil {
			continue
		}
		if key := fmt.Sprint(printer.Print(one.Name)); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.Directive
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneDirective(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneDirective attempts to merge all members of Directive into a singe *ast.Directive.
// If this cannot be done, this method will return an error.
func (m *Merger) OneDirective(curr []*ast.Directive, more ...*ast.Directive) (*ast.Directive, error) {
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
	var listArguments []*ast.Argument
	// range over the parent struct and collect properties
	for _, one := range all {
		listName = append(listName, one.Name)
		listArguments = append(listArguments, one.Arguments...)
	}

	var errSet error

	// merge properties

	one := ast.NewDirective(nil)
	if merged, err := m.OneName(listName); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Name = merged
	}
	if merged, err := m.SimilarArgument(listArguments); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Arguments = merged
	}

	return one, errSet

}
