package meta

import (
	"reflect"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/richardwilkes/toolbox/errs"
)

// IsConcrete returns true if `kind` is a concrete type,
// it returns an error if kind does not exist
func IsConcrete(kind string) (bool, reflect.Type, error) {
	return known.isConcrete(kind)
}

// IsInterface returns true if `kind` is an interface type,
// it returns an error if kind does not exist
func IsInterface(kind string) (bool, reflect.Type, error) {
	return known.isInterface(kind)
}

// DoesImplement returns true if the `concrete` kind
// implements the `ifc` kind
func DoesImplement(concrete, ifc string) (bool, error) {
	return known.doesImplement(concrete, ifc)
}

// HasFieldKind returns true if the `parentKind` kind is a concrete type
// and has a field named `fieldName` with the kind `fieldKind`
func HasFieldKind(parentKind, fieldKind, fieldName string) (bool, reflect.StructField, error) {
	return known.hasFieldKind(parentKind, fieldKind, fieldName)
}

// AllKinds returns all known types
func AllKinds() map[string]reflect.Type {
	return known.allKinds()
}

// AllConcrete returns all known concrete types,
// keyed by their `kind`
func AllConcrete() map[string]reflect.Type {
	return known.allConcrete()
}

// AllInterface returns all known interface types,
// keyed by their `kind`
func AllInterface() map[string]reflect.Type {
	return known.allInterface()
}

// AllFields returns all fields in `kind`,
// this will return an error of `kind` is not a known concrete type
func AllFields(kind string) ([]reflect.StructField, error) {
	return known.allFields(kind)
}

// AllImplementers finds all known types that implement `kind`,
// this will return an error if `kind` is not an interface type
func AllImplementers(kind string) (map[string]reflect.Type, error) {
	return known.allImplementers(kind)
}

// ----------------
// internal helpers

type allkinds map[string]reflect.Type

func (k allkinds) isConcrete(kind string) (bool, reflect.Type, error) {
	found, exists := k[kind]
	if !exists {
		return false, nil, errs.Newf("kind %q not found", kind)
	}
	return found.Kind() == reflect.Ptr, found, nil
}

func (k allkinds) isInterface(kind string) (bool, reflect.Type, error) {
	found, exists := k[kind]
	if !exists {
		return false, nil, errs.Newf("kind %q not found", kind)
	}
	return found.Kind() == reflect.Interface, found, nil
}

func (k allkinds) hasFieldKind(parentKind, fieldKind, fieldName string) (bool, reflect.StructField, error) {
	var field reflect.StructField

	isParentConcrete, parent, err := k.isConcrete(parentKind)
	if err != nil {
		return false, field, errs.Wrap(err)
	} else if !isParentConcrete {
		return false, field, errs.Newf("parentKind %q was not a concrete type", parentKind)
	}

	if field, found := parent.Elem().FieldByName(fieldName); found {
		child, ok := k[fieldKind]
		if !ok {
			return false, field, errs.Newf("fieldKind %q was not found", fieldKind)
		}
		return field.Type == child, field, nil
	}

	return false, field, nil
}

func (k allkinds) doesImplement(concrete, ifc string) (bool, error) {
	ok, foundIfc, err := k.isInterface(ifc)
	if err != nil {
		return false, errs.Wrap(err)
	}
	if !ok {
		return false, errs.Newf("`ifc` argument %q was not an interface... %q", ifc, foundIfc)
	}

	ok, foundConcrete, err := k.isConcrete(concrete)
	if err != nil {
		return false, errs.Wrap(err)
	}
	if !ok {
		return false, errs.Newf("`concrete` argument %q was not a concrete type... %q", concrete, foundConcrete)
	}

	implementers, err := k.allImplementers(ifc)
	if err != nil {
		return false, errs.Wrap(err)
	}

	if _, ok := implementers[concrete]; ok {
		return true, nil
	}
	return false, nil
}

func (k allkinds) allKinds() map[string]reflect.Type {
	return k
}

func (k allkinds) allConcrete() map[string]reflect.Type {
	all := make(map[string]reflect.Type)

	for kind, rtype := range k {
		if rtype.Kind() == reflect.Ptr {
			all[kind] = rtype
		}
	}

	return all
}

func (k allkinds) allInterface() map[string]reflect.Type {
	all := make(map[string]reflect.Type)

	for kind, rtype := range k {
		if rtype.Kind() == reflect.Interface {
			all[kind] = rtype
		}
	}

	return all
}

func (k allkinds) allFields(kind string) ([]reflect.StructField, error) {
	rtype, exists := k[kind]
	if !exists {
		return nil, errs.Newf("kind %q unknown", kind)
	}
	if rtk := rtype.Kind(); rtk != reflect.Ptr {
		return nil, errs.Newf("kind %q not concrete type, is %q", kind, rtk)
	}

	var fields []reflect.StructField

	elem := rtype.Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		switch field.Name {
		case "Kind", "Loc":
			continue
		default:
			fields = append(fields, field)
		}
	}

	return fields, nil
}

func (k allkinds) allImplementers(kind string) (map[string]reflect.Type, error) {
	rtype, exists := k[kind]
	if !exists {
		return nil, errs.Newf("kind %q unknown", kind)
	}
	if rtk := rtype.Kind(); rtk != reflect.Interface {
		return nil, errs.Newf("kind %q not interface type, is %q", kind, rtk)
	}

	all := make(map[string]reflect.Type)

	for ckind, ctype := range k.allConcrete() {
		if ctype.Implements(rtype) {
			all[ckind] = ctype
		}
	}

	return all, nil
}

var known = allkinds{
	"Node": reflect.TypeOf(new(ast.Node)).Elem(),

	// Name
	kinds.Name: reflect.TypeOf(ast.NewName(nil)),

	// Document
	kinds.Document:            reflect.TypeOf(ast.NewDocument(nil)),
	kinds.OperationDefinition: reflect.TypeOf(ast.NewOperationDefinition(nil)),
	kinds.VariableDefinition:  reflect.TypeOf(ast.NewVariableDefinition(nil)),
	kinds.Variable:            reflect.TypeOf(ast.NewVariable(nil)),
	kinds.SelectionSet:        reflect.TypeOf(ast.NewSelectionSet(nil)),
	kinds.Field:               reflect.TypeOf(ast.NewField(nil)),
	kinds.Argument:            reflect.TypeOf(ast.NewArgument(nil)),
	"Selection":               reflect.TypeOf(new(ast.Selection)).Elem(),

	// Fragments
	kinds.FragmentSpread:     reflect.TypeOf(ast.NewFragmentSpread(nil)),
	kinds.InlineFragment:     reflect.TypeOf(ast.NewInlineFragment(nil)),
	kinds.FragmentDefinition: reflect.TypeOf(ast.NewFragmentDefinition(nil)),

	// Values
	kinds.IntValue:     reflect.TypeOf(ast.NewIntValue(nil)),
	kinds.FloatValue:   reflect.TypeOf(ast.NewFloatValue(nil)),
	kinds.StringValue:  reflect.TypeOf(ast.NewStringValue(nil)),
	kinds.BooleanValue: reflect.TypeOf(ast.NewBooleanValue(nil)),
	kinds.EnumValue:    reflect.TypeOf(ast.NewEnumValue(nil)),
	kinds.ObjectValue:  reflect.TypeOf(ast.NewObjectValue(nil)),
	kinds.ObjectField:  reflect.TypeOf(ast.NewObjectField(nil)),
	"Value":            reflect.TypeOf(new(ast.Value)).Elem(),

	// Directives
	kinds.Directive: reflect.TypeOf(ast.NewDirective(nil)),

	// Types
	kinds.Named:   reflect.TypeOf(ast.NewNamed(nil)),
	kinds.List:    reflect.TypeOf(ast.NewList(nil)),
	kinds.NonNull: reflect.TypeOf(ast.NewNonNull(nil)),
	"Type":        reflect.TypeOf(new(ast.Type)).Elem(),

	// Type System Definitions
	kinds.SchemaDefinition:        reflect.TypeOf(ast.NewSchemaDefinition(nil)),
	kinds.OperationTypeDefinition: reflect.TypeOf(ast.NewOperationTypeDefinition(nil)),

	// Types Definitions
	kinds.ScalarDefinition:      reflect.TypeOf(ast.NewScalarDefinition(nil)),
	kinds.ObjectDefinition:      reflect.TypeOf(ast.NewObjectDefinition(nil)),
	kinds.FieldDefinition:       reflect.TypeOf(ast.NewFieldDefinition(nil)),
	kinds.InputValueDefinition:  reflect.TypeOf(ast.NewInputValueDefinition(nil)),
	kinds.InterfaceDefinition:   reflect.TypeOf(ast.NewInterfaceDefinition(nil)),
	kinds.UnionDefinition:       reflect.TypeOf(ast.NewUnionDefinition(nil)),
	kinds.EnumDefinition:        reflect.TypeOf(ast.NewEnumDefinition(nil)),
	kinds.EnumValueDefinition:   reflect.TypeOf(ast.NewEnumValueDefinition(nil)),
	kinds.InputObjectDefinition: reflect.TypeOf(ast.NewInputObjectDefinition(nil)),

	// Types Extensions
	kinds.TypeExtensionDefinition: reflect.TypeOf(ast.NewTypeExtensionDefinition(nil)),

	// Directive Definitions
	kinds.DirectiveDefinition: reflect.TypeOf(ast.NewDirectiveDefinition(nil)),
}
