package directives

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/richardwilkes/toolbox/errs"
)

// GetDirectives returns nil if the node does not have directives
func GetDirectives(node ast.Node) []*ast.Directive {
	switch one := node.(type) {
	case *ast.ScalarDefinition:
		return one.Directives
	case *ast.FieldDefinition:
		return one.Directives
	case *ast.EnumDefinition:
		return one.Directives
	case *ast.EnumValueDefinition:
		return one.Directives
	case *ast.ObjectDefinition:
		return one.Directives
	case *ast.TypeExtensionDefinition:
		return one.Definition.Directives
	case *ast.InterfaceDefinition:
		return one.Directives
	}
	return nil
}

// SetDirectives returns false if the node does not accept directives
func SetDirectives(node ast.Node, directives []*ast.Directive) bool {
	switch one := node.(type) {
	case *ast.ScalarDefinition:
		one.Directives = directives
	case *ast.FieldDefinition:
		one.Directives = directives
	case *ast.EnumDefinition:
		one.Directives = directives
	case *ast.EnumValueDefinition:
		one.Directives = directives
	case *ast.ObjectDefinition:
		one.Directives = directives
	case *ast.TypeExtensionDefinition:
		one.Definition.Directives = directives
	case *ast.InterfaceDefinition:
		one.Directives = directives
	default:
		return false
	}

	return true
}

// GetDirective gets a directive by name
func GetDirective(node ast.Node, dirName string) *ast.Directive {
	for _, one := range GetDirectives(node) {
		if one.Name.Value == dirName {
			return one
		}
	}
	return nil

}

// GetDirectiveArg returns nil if the argument was not found
func GetDirectiveArg(node ast.Node, dirName, argName string) ast.Value {
	dir := GetDirective(node, dirName)
	if dir == nil {
		return nil
	}

	for _, arg := range dir.Arguments {
		if arg.Name.Value == argName {
			return arg.Value
		}
	}

	return nil
}

// GetDirectiveArgEnum returns `false` if directive or argument was not found,
// or returns an `error` if it was found but it was not a string
func GetDirectiveArgEnum(node ast.Node, dirName, argName string) (string, bool, error) {
	val := GetDirectiveArg(node, dirName, argName)
	if val == nil {
		return "", false, nil
	}

	enumVal, ok := val.(*ast.EnumValue)
	if !ok {
		return "", false, errs.Newf("@%s(%s: ...) not a string, is %T", dirName, argName, val)
	}
	if enumVal == nil {
		return "", false, errs.Newf("@%s(%s: ...) was a string, but is nil", dirName, argName)
	}

	return enumVal.Value, true, nil
}

// GetDirectiveArgStr returns `false` if directive or argument was not found,
// or returns an `error` if it was found but it was not a string
func GetDirectiveArgStr(node ast.Node, dirName, argName string) (string, bool, error) {
	val := GetDirectiveArg(node, dirName, argName)
	if val == nil {
		return "", false, nil
	}

	strVal, ok := val.(*ast.StringValue)
	if !ok {
		return "", false, errs.Newf("@%s(%s: ...) not a string, is %T", dirName, argName, val)
	}
	if strVal == nil {
		return "", false, errs.Newf("@%s(%s: ...) was a string, but is nil", dirName, argName)
	}

	return strVal.Value, true, nil
}

// GetDirectiveArgBool returns `(false, false, nil)` if directive or argument was not found,
// or returns an `error` if it was found but it was not a boolean
func GetDirectiveArgBool(node ast.Node, dirName, argName string) (bool, bool, error) {
	val := GetDirectiveArg(node, dirName, argName)
	if val == nil {
		return false, false, nil
	}

	boolVal, ok := val.(*ast.BooleanValue)
	if !ok {
		return false, false, errs.Newf("@%s(%s: ...) not a boolean, is %T", dirName, argName, val)
	}
	if boolVal == nil {
		return false, false, errs.Newf("@%s(%s: ...) was a boolean, but is nil", dirName, argName)
	}

	return boolVal.Value, true, nil
}
