package mergers

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
	"github.com/richardwilkes/toolbox/errs"
)

// Merger ...
type Merger struct {
	getObjValueID func(objval *ast.ObjectValue) string
}

// NewMerger ...
func NewMerger(opts ...func(*Merger) error) (*Merger, error) {
	m := &Merger{
		getObjValueID: defaultObjValueID,
	}

	for _, one := range opts {
		if err := one(m); err != nil {
			return nil, errs.Wrap(err)
		}
	}

	return m, nil
}

// IsSameValue detects if two values are the same
func (m *Merger) IsSameValue(rawLeft, rawRight ast.Value) (bool, error) {
	switch left := rawLeft.(type) {
	// always merge lists
	case *ast.ListValue:
		return true, nil

	// simple types can have their direct equality checked
	case *ast.BooleanValue,
		*ast.EnumValue,
		*ast.IntValue,
		*ast.FloatValue,
		*ast.StringValue:
		return rawLeft.GetValue() == rawRight.GetValue(), nil

	// objects types must have external key detection
	case *ast.ObjectValue:
		if right, ok := rawRight.(*ast.ObjectValue); ok {
			if hasLeft, hasRight := left == nil, right == nil; hasLeft == hasRight {
				return true, nil
			} else if hasLeft != hasRight {
				return false, nil
			}
			leftID := m.getObjValueID(left)
			rightID := m.getObjValueID(right)
			return leftID == rightID, nil
		}
	}

	return false, errs.Newf("cannot compare ast.Value of types left: %T ... right: %T", rawLeft, rawRight)
}

// ---
func (m *Merger) SimilarNodes(curr []ast.Node, more ...ast.Node) ([]ast.Node, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneNode(curr []ast.Node, more ...ast.Node) (ast.Node, error) {
	return nil, errs.Newf("unimplemented")
}

// SimilarScalars ...
func (m *Merger) SimilarScalars(curr []*ast.ScalarDefinition, more ...*ast.ScalarDefinition) ([]*ast.ScalarDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.ScalarDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.ScalarDefinition
	for _, group := range all {
		if merged, err := m.OneScalar(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarEnums ...
func (m *Merger) SimilarEnums(curr []*ast.EnumDefinition, more ...*ast.EnumDefinition) ([]*ast.EnumDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.EnumDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.EnumDefinition
	for _, group := range all {
		if merged, err := m.OneEnum(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarEnumValues ...
func (m *Merger) SimilarEnumValues(curr []*ast.EnumValueDefinition, more ...*ast.EnumValueDefinition) ([]*ast.EnumValueDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.EnumValueDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.EnumValueDefinition
	for _, group := range all {
		if merged, err := m.OneEnumValue(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarObjects ...
func (m *Merger) SimilarObjects(curr []*ast.ObjectDefinition, more ...*ast.ObjectDefinition) ([]*ast.ObjectDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.ObjectDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.ObjectDefinition
	for _, group := range all {
		if merged, err := m.OneObject(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarInputObjects ...
func (m *Merger) SimilarInputObjects(curr []*ast.InputObjectDefinition, more ...*ast.InputObjectDefinition) ([]*ast.InputObjectDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.InputObjectDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.InputObjectDefinition
	for _, group := range all {
		if merged, err := m.OneInputObject(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarUnions ...
func (m *Merger) SimilarUnions(curr []*ast.UnionDefinition, more ...*ast.UnionDefinition) ([]*ast.UnionDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.UnionDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.UnionDefinition
	for _, group := range all {
		if merged, err := m.OneUnion(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarExtensions ...
func (m *Merger) SimilarExtensions(curr []*ast.TypeExtensionDefinition, more ...*ast.TypeExtensionDefinition) ([]*ast.TypeExtensionDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.TypeExtensionDefinition)
	for _, one := range set {
		curr, _ := all[one.Definition.Name.Value]
		all[one.Definition.Name.Value] = append(curr, one)
	}

	var out []*ast.TypeExtensionDefinition
	for _, group := range all {
		if merged, err := m.OneExtension(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarInterfaces ...
func (m *Merger) SimilarInterfaces(curr []*ast.InterfaceDefinition, more ...*ast.InterfaceDefinition) ([]*ast.InterfaceDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.InterfaceDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.InterfaceDefinition
	for _, group := range all {
		if merged, err := m.OneInterface(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarDirectives ...
func (m *Merger) SimilarDirectives(curr []*ast.Directive, more ...*ast.Directive) ([]*ast.Directive, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.Directive)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.Directive
	for _, group := range all {
		if merged, err := m.OneDirective(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarDirectiveDefinitions ...
func (m *Merger) SimilarDirectiveDefinitions(curr []*ast.DirectiveDefinition, more ...*ast.DirectiveDefinition) ([]*ast.DirectiveDefinition, error) {
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.DirectiveDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.DirectiveDefinition
	for _, group := range all {
		if merged, err := m.OneDirectiveDefinition(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarInputObjects ...
func (m *Merger) SimilarFields(curr []*ast.FieldDefinition, more ...*ast.FieldDefinition) ([]*ast.FieldDefinition, error) {
	return nil, errs.Newf("unimplemented")
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.EnumValueDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.EnumValueDefinition
	for _, group := range all {
		if merged, err := m.OneEnumValue(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarInputObjects ...
func (m *Merger) SimilarArguments(curr []*ast.Argument, more ...*ast.Argument) ([]*ast.Argument, error) {
	return nil, errs.Newf("unimplemented")
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.EnumValueDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.EnumValueDefinition
	for _, group := range all {
		if merged, err := m.OneEnumValue(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarInputObjects ...
func (m *Merger) SimilarValues(curr []ast.Value, more ...ast.Value) ([]ast.Value, error) {
	return nil, errs.Newf("unimplemented")
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.EnumValueDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.EnumValueDefinition
	for _, group := range all {
		if merged, err := m.OneEnumValue(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}

// SimilarInputObjects ...
func (m *Merger) SimilarInputValues(curr []*ast.InputValueDefinition, more ...ast.InputValueDefinition) ([]*ast.InputValueDefinition, error) {
	return nil, errs.Newf("unimplemented")
	set := append(curr, more...)
	switch len(set) {
	case 0:
		return nil, nil
	case 1:
		return set, nil
	}

	all := make(map[string][]*ast.EnumValueDefinition)
	for _, one := range set {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.EnumValueDefinition
	for _, group := range all {
		if merged, err := m.OneEnumValue(group); err != nil {
			return nil, errs.Wrap(err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, nil
}
func (m *Merger) OneScalar(curr []*ast.ScalarDefinition, more ...*ast.ScalarDefinition) (*ast.ScalarDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneEnum(curr []*ast.EnumDefinition, more ...*ast.EnumDefinition) (*ast.EnumDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneEnumValue(curr []*ast.EnumValueDefinition, more ...*ast.EnumValueDefinition) (*ast.EnumValueDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneObject(curr []*ast.ObjectDefinition, more ...*ast.ObjectDefinition) (*ast.ObjectDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneInputObject(curr []*ast.InputObjectDefinition, more ...*ast.InputObjectDefinition) (*ast.InputObjectDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneUnion(curr []*ast.UnionDefinition, more ...*ast.UnionDefinition) (*ast.UnionDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneExtension(curr []*ast.TypeExtensionDefinition, more ...*ast.TypeExtensionDefinition) (*ast.TypeExtensionDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneInterface(curr []*ast.InterfaceDefinition, more ...*ast.InterfaceDefinition) (*ast.InterfaceDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneDirective(curr []*ast.Directive, more ...*ast.EnumValueDefinition) (*ast.Directive, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneDirectiveDefinition(curr []*ast.DirectiveDefinition, more ...*ast.DirectiveDefinition) (*ast.DirectiveDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneField(curr []*ast.FieldDefinition, more ...*ast.FieldDefinition) (*ast.FieldDefinition, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneArgument(curr []*ast.Argument, more ...*ast.Argument) (*ast.Argument, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneValue(curr []ast.Value, more ...ast.Value) (ast.Value, error) {
	return nil, errs.Newf("unimplemented")
}
func (m *Merger) OneInputValue(curr []*ast.InputValueDefinition, more ...ast.InputValueDefinition) (*ast.InputValueDefinition, error) {
	return nil, errs.Newf("unimplemented")
}

// --------------------

// ---------------------
// helpers and constants

var defaultObjValueID = func(objval *ast.ObjectValue) string {
	if objval == nil {
		return ""
	}
	return fmt.Sprint(printer.Print(objval))
}
