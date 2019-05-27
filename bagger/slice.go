package bagger

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/richardwilkes/toolbox/errs"
)

// AsNodes takes a map or slice of the pointer type of the node
// and turns it into an array of the interface `ast.Node`
// to make passing it to generic functions easier
func AsNodes(raw interface{}) ([]ast.Node, error) {
	var all []ast.Node

	switch src := raw.(type) {
	case map[string]*ast.ScalarDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.EnumDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.ObjectDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.InputObjectDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.UnionDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.InterfaceDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.TypeExtensionDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.DirectiveDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case map[string]*ast.FieldDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.ScalarDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.EnumDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.ObjectDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.InputObjectDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.UnionDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.InterfaceDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.TypeExtensionDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.DirectiveDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	case []*ast.FieldDefinition:
		for _, one := range src {
			all = append(all, one)
		}
	default:
		return nil, errs.Newf("%T cannot be converted to []ast.Node", src)
	}

	return all, nil
}
