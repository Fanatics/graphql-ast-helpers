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

// SimilarDocument merges declarations of Document that share the same Document value.
// This uses the default basic merge strategy.
func SimilarDocument(curr []*ast.Document, more ...*ast.Document) ([]*ast.Document, error) {
	return Basic.SimilarDocument(curr, more...)
}

// OneDocument attempts to merge all members of Document into a singe *ast.Document.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneDocument(curr []*ast.Document, more ...*ast.Document) (*ast.Document, error) {
	return Basic.OneDocument(curr, more...)
}

// SimilarDocument merges declarations of Document that share the same Document value.
func (m *Merger) SimilarDocument(curr []*ast.Document, more ...*ast.Document) ([]*ast.Document, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]*ast.Document)
	for _, one := range all {
		if one == nil {
			continue
		}
		if key := "document"; key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []*ast.Document
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneDocument(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneDocument attempts to merge all members of Document into a singe *ast.Document.
// If this cannot be done, this method will return an error.
func (m *Merger) OneDocument(curr []*ast.Document, more ...*ast.Document) (*ast.Document, error) {
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
	var listDefinitions []ast.Node
	// range over the parent struct and collect properties
	for _, one := range all {
		listDefinitions = append(listDefinitions, one.Definitions...)
	}

	var errSet error

	// merge properties

	one := ast.NewDocument(nil)
	if merged, err := m.SimilarNode(listDefinitions); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.Definitions = merged
	}

	return one, errSet

}
