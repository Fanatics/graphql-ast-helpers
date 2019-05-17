package creates

import (
	"regexp"

	"github.com/graphql-go/graphql/language/ast"
)

// Name ...
func Name(name string) *ast.Name {
	return ast.NewName(&ast.Name{
		Value: name,
	})
}

// NamedType ...
func NamedType(name string) *ast.Named {
	return ast.NewNamed(&ast.Named{
		Name: Name(name),
	})
}

var matchRawType = regexp.MustCompile(`(\[)?([_A-Za-z][_0-9A-Za-z]*)(\!)?(\])?(\!)?`)

// Type parses raw like "[String!]!" into ast.Type
// If the type is unparsable, nil is returned
func Type(raw string) ast.Type {
	match := matchRawType.FindAllStringSubmatch(raw, -1)
	if match == nil {
		return nil
	}

	var (
		tokenOpenList     = match[0][1]
		tokenName         = match[0][2]
		tokenRequired     = match[0][3]
		tokenCloseList    = match[0][4]
		tokenListRequired = match[0][5]
	)

	// an unmatched open or close
	if len(tokenOpenList) != len(tokenCloseList) {
		return nil
	}

	var node ast.Type = NamedType(tokenName)

	if tokenRequired == "!" {
		node = AsNonNull(node)
	}
	if tokenOpenList == "[" && tokenCloseList == "]" {
		node = AsList(node)
		if tokenListRequired == "!" {
			node = AsNonNull(node)
		}
	}

	return node
}

// AsNonNull ...
func AsNonNull(kind ast.Type) *ast.NonNull {
	return ast.NewNonNull(&ast.NonNull{
		Type: kind,
	})
}

// AsList ...
func AsList(kind ast.Type) *ast.List {
	return ast.NewList(&ast.List{
		Type: kind,
	})
}

// ValueString ...
func ValueString(value string) *ast.StringValue {
	return ast.NewStringValue(&ast.StringValue{
		Value: value,
	})
}

// ValueBoolean ...
func ValueBoolean(value bool) *ast.BooleanValue {
	return ast.NewBooleanValue(&ast.BooleanValue{
		Value: value,
	})
}

// ValueList ...
func ValueList(vals ...ast.Value) *ast.ListValue {
	return ast.NewListValue(&ast.ListValue{
		Values: vals,
	})
}

// ObjVal ...
func ObjVal(fields ...*ast.ObjectField) *ast.ObjectValue {
	return ast.NewObjectValue(&ast.ObjectValue{
		Fields: fields,
	})
}

// ObjValField ...
func ObjValField(name string, val ast.Value) *ast.ObjectField {
	return ast.NewObjectField(&ast.ObjectField{
		Name:  Name(name),
		Value: val,
	})
}

// Arg ...
func Arg(name string, value ast.Value) *ast.Argument {
	return ast.NewArgument(&ast.Argument{
		Name:  Name(name),
		Value: value,
	})
}

// ArgString ...
func ArgString(name string, value string) *ast.Argument {
	return Arg(name, ValueString(value))
}

// ArgBoolean ...
func ArgBoolean(name string, value bool) *ast.Argument {
	return Arg(name, ValueBoolean(value))
}

// ArgEnum ...
func ArgEnum(name string, value string) *ast.Argument {
	return Arg(name, ast.NewEnumValue(&ast.EnumValue{
		Value: value,
	}))
}

// ObjField ...
func ObjField(name string) *ast.ObjectField {
	return ast.NewObjectField(&ast.ObjectField{
		Name: Name(name),
	})
}

// InputVal ...
func InputVal(name string, kind ast.Type) *ast.InputValueDefinition {
	return ast.NewInputValueDefinition(&ast.InputValueDefinition{
		Name: Name(name),
		Type: kind,
	})
}

// Scalar ...
func Scalar(name string) *ast.ScalarDefinition {
	return ast.NewScalarDefinition(&ast.ScalarDefinition{
		Name:        Name(name),
		Description: ValueString(""),
	})
}

// EnumVal ...
func EnumVal(name string) *ast.EnumValueDefinition {
	return ast.NewEnumValueDefinition(&ast.EnumValueDefinition{
		Name: Name(name),
	})
}

// Enum ...
func Enum(name string, vals []string) *ast.EnumDefinition {
	enum := ast.NewEnumDefinition(&ast.EnumDefinition{
		Name: Name(name),
	})
	for _, val := range vals {
		enum.Values = append(enum.Values, EnumVal(val))
	}
	return enum
}

// Interface ...
func Interface(name string) *ast.InterfaceDefinition {
	return ast.NewInterfaceDefinition(&ast.InterfaceDefinition{
		Name:        Name(name),
		Description: ValueString(""),
	})
}

// Union ...
func Union(name string, types ...string) *ast.UnionDefinition {
	var typeNodes []*ast.Named
	for _, one := range types {
		typeNodes = append(typeNodes, NamedType(one))
	}
	return ast.NewUnionDefinition(&ast.UnionDefinition{
		Name:        Name(name),
		Description: ValueString(""),
		Types:       typeNodes,
	})
}

// Obj ...
func Obj(name string) *ast.ObjectDefinition {
	return ast.NewObjectDefinition(&ast.ObjectDefinition{
		Name:        Name(name),
		Description: ValueString(""),
	})
}

// InputObj ...
func InputObj(name string) *ast.InputObjectDefinition {
	return ast.NewInputObjectDefinition(&ast.InputObjectDefinition{
		Name:        Name(name),
		Description: ValueString(""),
	})
}

// ObjExt ...
func ObjExt(name string) *ast.TypeExtensionDefinition {
	return ast.NewTypeExtensionDefinition(&ast.TypeExtensionDefinition{
		Definition: Obj(name),
	})
}

// Field ...
func Field(name string, kind ast.Type) *ast.FieldDefinition {
	return ast.NewFieldDefinition(&ast.FieldDefinition{
		Name:        Name(name),
		Type:        kind,
		Description: ValueString(""),
	})
}

// FieldFrom attemps to parse the raw type into a proper type
func FieldFrom(name string, rawType string) *ast.FieldDefinition {
	return ast.NewFieldDefinition(&ast.FieldDefinition{
		Name:        Name(name),
		Type:        Type(rawType),
		Description: ValueString(""),
	})
}

// FieldByFlat creates a field with the appropriate type
// wrapping described by a set of flags
func FieldByFlat(name string, typeName string, isRequired bool, isMany bool, isManyRequired bool) *ast.FieldDefinition {
	var kind ast.Type = NamedType(typeName)
	if isRequired {
		kind = AsNonNull(kind)
	}
	if isMany {
		kind = AsList(kind)
		if isManyRequired {
			kind = AsNonNull(kind)
		}
	}
	return Field(name, kind)
}

// Ext ...
func Ext(name string) *ast.TypeExtensionDefinition {
	return ast.NewTypeExtensionDefinition(&ast.TypeExtensionDefinition{
		Definition: Obj(name),
	})
}

// DirectiveDef ...
func DirectiveDef(name string, vals ...*ast.InputValueDefinition) *ast.DirectiveDefinition {
	return ast.NewDirectiveDefinition(&ast.DirectiveDefinition{
		Name:      Name(name),
		Arguments: vals,
	})
}

// Directive ...
func Directive(name string, args ...*ast.Argument) *ast.Directive {
	return ast.NewDirective(&ast.Directive{
		Name:      Name(name),
		Arguments: args,
	})
}

// Query creates the `query` field for the top-level `schema` node
func Query(objectName string) *ast.OperationTypeDefinition {
	return ast.NewOperationTypeDefinition(&ast.OperationTypeDefinition{
		Operation: ast.OperationTypeQuery,
		Type:      NamedType(objectName),
	})
}

// Mutation creates the `mutation` field for the top-level `schema` node
func Mutation(objectName string) *ast.OperationTypeDefinition {
	return ast.NewOperationTypeDefinition(&ast.OperationTypeDefinition{
		Operation: ast.OperationTypeMutation,
		Type:      NamedType(objectName),
	})
}

// Subscription creates the `subscription` field for the top-level `schema` node
func Subscription(objectName string) *ast.OperationTypeDefinition {
	return ast.NewOperationTypeDefinition(&ast.OperationTypeDefinition{
		Operation: ast.OperationTypeSubscription,
		Type:      NamedType(objectName),
	})
}

// Schema creates the top level `schema` node. Set argument strings to `""` to
// ignore a given type
func Schema(queryName, mutationName, subscriptionName string) *ast.SchemaDefinition {
	schema := ast.NewSchemaDefinition(nil)

	if queryName != "" {
		schema.OperationTypes = append(
			schema.OperationTypes,
			Query(queryName),
		)
	}
	if mutationName != "" {
		schema.OperationTypes = append(
			schema.OperationTypes,
			Mutation(mutationName),
		)
	}
	if subscriptionName != "" {
		schema.OperationTypes = append(
			schema.OperationTypes,
			Subscription(subscriptionName),
		)
	}

	return schema
}
