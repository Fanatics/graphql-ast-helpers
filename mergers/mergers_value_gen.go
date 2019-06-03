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

// SimilarValue merges declarations of Value that share the same Value value.
func (m *Merger) SimilarValue(curr []ast.Value, more ...ast.Value) ([]ast.Value, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]ast.Value)
	for _, one := range all {
		if key := m.getValueID(one); key != "" {
			curr, _ := groups[key]
			groups[key] = append(curr, one)
		}
	}

	var out []ast.Value
	var errSet error

	for _, group := range groups {
		if merged, err := m.OneValue(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// OneValue attempts to merge all members of Value into a singe ast.Value.
// If this cannot be done, this method will return an error.
func (m *Merger) OneValue(curr []ast.Value, more ...ast.Value) (ast.Value, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}

	var errSet error

	// merge properties

	switch all[0].(type) {
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
	default:
		errSet = errs.Append(errSet, errs.Newf("type %T unknown", all[0]))
	}

	return nil, errSet

}
