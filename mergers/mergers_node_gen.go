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

// SimilarNode merges declarations of Node that share the same Node value.
// This uses the default basic merge strategy.
func SimilarNode(curr []ast.Node, more ...ast.Node) ([]ast.Node, error) {
	return Basic.SimilarNode(curr, more...)
}

// OneNode attempts to merge all members of Node into a singe ast.Node.
// If this cannot be done, this method will return an error.
// This uses the default basic merge strategy.
func OneNode(curr []ast.Node, more ...ast.Node) (ast.Node, error) {
	return Basic.OneNode(curr, more...)
}

// SimilarNode merges declarations of Node that share the same Node value.
func (m *Merger) SimilarNode(curr []ast.Node, more ...ast.Node) ([]ast.Node, error) {
	if m == nil {
		return nil, errs.New("merger strategy was nil")
	}

	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]ast.Node)
	for _, one := range all {
		if key := m.getNodeID(one); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []ast.Node
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneNode(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneNode attempts to merge all members of Node into a singe ast.Node.
// If this cannot be done, this method will return an error.
func (m *Merger) OneNode(curr []ast.Node, more ...ast.Node) (ast.Node, error) {
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

	var errSet error

	// merge properties

	switch all[0].(type) {
	case *ast.Argument:
		var set []*ast.Argument
		for _, single := range all {
			v, ok := single.(*ast.Argument)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.Argument but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneArgument(set)
	case *ast.BooleanValue:
		var set []*ast.BooleanValue
		for _, single := range all {
			v, ok := single.(*ast.BooleanValue)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.BooleanValue but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneBooleanValue(set)
	case *ast.Directive:
		var set []*ast.Directive
		for _, single := range all {
			v, ok := single.(*ast.Directive)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.Directive but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneDirective(set)
	case *ast.DirectiveDefinition:
		var set []*ast.DirectiveDefinition
		for _, single := range all {
			v, ok := single.(*ast.DirectiveDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.DirectiveDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneDirectiveDefinition(set)
	case *ast.Document:
		var set []*ast.Document
		for _, single := range all {
			v, ok := single.(*ast.Document)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.Document but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneDocument(set)
	case *ast.EnumDefinition:
		var set []*ast.EnumDefinition
		for _, single := range all {
			v, ok := single.(*ast.EnumDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.EnumDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneEnumDefinition(set)
	case *ast.EnumValue:
		var set []*ast.EnumValue
		for _, single := range all {
			v, ok := single.(*ast.EnumValue)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.EnumValue but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneEnumValue(set)
	case *ast.EnumValueDefinition:
		var set []*ast.EnumValueDefinition
		for _, single := range all {
			v, ok := single.(*ast.EnumValueDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.EnumValueDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneEnumValueDefinition(set)
	case *ast.Field:
		var set []*ast.Field
		for _, single := range all {
			v, ok := single.(*ast.Field)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.Field but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneField(set)
	case *ast.FieldDefinition:
		var set []*ast.FieldDefinition
		for _, single := range all {
			v, ok := single.(*ast.FieldDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.FieldDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneFieldDefinition(set)
	case *ast.FloatValue:
		var set []*ast.FloatValue
		for _, single := range all {
			v, ok := single.(*ast.FloatValue)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.FloatValue but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneFloatValue(set)
	case *ast.FragmentDefinition:
		var set []*ast.FragmentDefinition
		for _, single := range all {
			v, ok := single.(*ast.FragmentDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.FragmentDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneFragmentDefinition(set)
	case *ast.FragmentSpread:
		var set []*ast.FragmentSpread
		for _, single := range all {
			v, ok := single.(*ast.FragmentSpread)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.FragmentSpread but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneFragmentSpread(set)
	case *ast.InlineFragment:
		var set []*ast.InlineFragment
		for _, single := range all {
			v, ok := single.(*ast.InlineFragment)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.InlineFragment but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneInlineFragment(set)
	case *ast.InputObjectDefinition:
		var set []*ast.InputObjectDefinition
		for _, single := range all {
			v, ok := single.(*ast.InputObjectDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.InputObjectDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneInputObjectDefinition(set)
	case *ast.InputValueDefinition:
		var set []*ast.InputValueDefinition
		for _, single := range all {
			v, ok := single.(*ast.InputValueDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.InputValueDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneInputValueDefinition(set)
	case *ast.IntValue:
		var set []*ast.IntValue
		for _, single := range all {
			v, ok := single.(*ast.IntValue)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.IntValue but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneIntValue(set)
	case *ast.InterfaceDefinition:
		var set []*ast.InterfaceDefinition
		for _, single := range all {
			v, ok := single.(*ast.InterfaceDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.InterfaceDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneInterfaceDefinition(set)
	case *ast.List:
		var set []*ast.List
		for _, single := range all {
			v, ok := single.(*ast.List)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.List but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneList(set)
	case *ast.Name:
		var set []*ast.Name
		for _, single := range all {
			v, ok := single.(*ast.Name)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.Name but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneName(set)
	case *ast.Named:
		var set []*ast.Named
		for _, single := range all {
			v, ok := single.(*ast.Named)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.Named but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneNamed(set)
	case *ast.NonNull:
		var set []*ast.NonNull
		for _, single := range all {
			v, ok := single.(*ast.NonNull)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.NonNull but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneNonNull(set)
	case *ast.ObjectDefinition:
		var set []*ast.ObjectDefinition
		for _, single := range all {
			v, ok := single.(*ast.ObjectDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.ObjectDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneObjectDefinition(set)
	case *ast.ObjectField:
		var set []*ast.ObjectField
		for _, single := range all {
			v, ok := single.(*ast.ObjectField)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.ObjectField but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneObjectField(set)
	case *ast.ObjectValue:
		var set []*ast.ObjectValue
		for _, single := range all {
			v, ok := single.(*ast.ObjectValue)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.ObjectValue but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneObjectValue(set)
	case *ast.OperationDefinition:
		var set []*ast.OperationDefinition
		for _, single := range all {
			v, ok := single.(*ast.OperationDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.OperationDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneOperationDefinition(set)
	case *ast.OperationTypeDefinition:
		var set []*ast.OperationTypeDefinition
		for _, single := range all {
			v, ok := single.(*ast.OperationTypeDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.OperationTypeDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneOperationTypeDefinition(set)
	case *ast.ScalarDefinition:
		var set []*ast.ScalarDefinition
		for _, single := range all {
			v, ok := single.(*ast.ScalarDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.ScalarDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneScalarDefinition(set)
	case *ast.SchemaDefinition:
		var set []*ast.SchemaDefinition
		for _, single := range all {
			v, ok := single.(*ast.SchemaDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.SchemaDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneSchemaDefinition(set)
	case *ast.SelectionSet:
		var set []*ast.SelectionSet
		for _, single := range all {
			v, ok := single.(*ast.SelectionSet)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.SelectionSet but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneSelectionSet(set)
	case *ast.StringValue:
		var set []*ast.StringValue
		for _, single := range all {
			v, ok := single.(*ast.StringValue)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.StringValue but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneStringValue(set)
	case *ast.TypeExtensionDefinition:
		var set []*ast.TypeExtensionDefinition
		for _, single := range all {
			v, ok := single.(*ast.TypeExtensionDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.TypeExtensionDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneTypeExtensionDefinition(set)
	case *ast.UnionDefinition:
		var set []*ast.UnionDefinition
		for _, single := range all {
			v, ok := single.(*ast.UnionDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.UnionDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneUnionDefinition(set)
	case *ast.Variable:
		var set []*ast.Variable
		for _, single := range all {
			v, ok := single.(*ast.Variable)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.Variable but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneVariable(set)
	case *ast.VariableDefinition:
		var set []*ast.VariableDefinition
		for _, single := range all {
			v, ok := single.(*ast.VariableDefinition)
			if !ok {
				errSet = errs.Append(errSet, errs.Newf("want *ast.VariableDefinition but got type %T", single))
				continue
			}
			set = append(set, v)
		}
		return m.OneVariableDefinition(set)
	default:
		errSet = errs.Append(errSet, errs.Newf("type %T unknown", all[0]))
	}

	return nil, errSet

}
