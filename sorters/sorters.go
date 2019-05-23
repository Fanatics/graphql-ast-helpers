package sorters

import (
	"sort"

	"github.com/graphql-go/graphql/language/ast"
)

// -------
// scalars

// SortScalars sorts an array of scalars in-place by name.
func SortScalars(scalars []*ast.ScalarDefinition) {
	sort.Slice(scalars, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			scalars[i].Name.Value,
			scalars[j].Name.Value,
		})
	})
}

// SortScalarsMap will sort a map of scalars by name and return a newly created slice.
func SortScalarsMap(all map[string]*ast.ScalarDefinition) []*ast.ScalarDefinition {
	var scalars []*ast.ScalarDefinition
	for _, one := range all {
		scalars = append(scalars, one)
	}
	SortScalars(scalars)
	return scalars
}

// -----
// enums

// SortEnums sorts an array of enums in-place by name,
// and also all of each enum's values by name.
func SortEnums(enums []*ast.EnumDefinition) {
	sort.Slice(enums, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			enums[i].Name.Value,
			enums[j].Name.Value,
		})
	})
	for _, enum := range enums {
		SortEnumValues(enum.Values)
	}
}

// SortEnumValues sorts an an array of enum values in-place by name.
func SortEnumValues(enumValues []*ast.EnumValueDefinition) {
	sort.Slice(enumValues, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			enumValues[i].Name.Value,
			enumValues[j].Name.Value,
		})
	})
}

// SortEnumsMap will sort a map of enums by name and return a newly created slice.
// This also sorts each enum's values by name.
func SortEnumsMap(all map[string]*ast.EnumDefinition) []*ast.EnumDefinition {
	var enums []*ast.EnumDefinition
	for _, one := range all {
		enums = append(enums, one)
	}
	SortEnums(enums)
	return enums
}

// ------
// fields

// SortFields sorts an array of fields in-place by name.
func SortFields(nodes []*ast.FieldDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// SortFieldsMap will sort a map of fields by name and return a newly created slice.
func SortFieldsMap(nodes []*ast.FieldDefinition) []*ast.FieldDefinition {
	var fields []*ast.FieldDefinition
	for _, one := range nodes {
		fields = append(fields, one)
	}
	SortFields(fields)
	return fields
}

// ------------
// object types

// SortObjects sorts an an array of objects in-place by name.
// This does not also sort fields, use `SortObjectFields` to get sorted fields.
func SortObjects(nodes []*ast.ObjectDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// SortObjectsMap will sort a map of objects by name and return a newly created slice.
// This will not sort the fields by name, use `SortObjectFields` to get sorted fields.
func SortObjectsMap(all map[string]*ast.ObjectDefinition) []*ast.ObjectDefinition {
	var nodes []*ast.ObjectDefinition
	for _, one := range all {
		nodes = append(nodes, one)
	}
	SortObjects(nodes)
	return nodes
}

// SortObjectsFields sorts the fields of each object in-place.
func SortObjectsFields(objs ...*ast.ObjectDefinition) {
	for _, one := range objs {
		SortFields(one.Fields)
	}
}

// -------------
// input objects

// SortInputObjects sorts an array of input objects in-place by name.
// This does not sort each input object's input values, use `SortInputObjectsFields` to sort input values.
func SortInputObjects(nodes []*ast.InputObjectDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// SortInputObjectsMap will sort a map of objects by name and return a newly created slice.
// This does not sort each input object's input values, use `SortInputObjectsFields` to sort input values.
func SortInputObjectsMap(all map[string]*ast.InputObjectDefinition) []*ast.InputObjectDefinition {
	var nodes []*ast.InputObjectDefinition
	for _, one := range all {
		nodes = append(nodes, one)
	}
	SortInputObjects(nodes)
	return nodes
}

// SortInputObjectsFields will sort each input object's fields in-place by name.
func SortInputObjectsFields(nodes ...*ast.InputObjectDefinition) {
	for _, one := range nodes {
		SortInputValues(one.Fields)
	}
}

// ------------
// input values

// SortInputValues sorts an array of input values in-place by name.
func SortInputValues(nodes []*ast.InputValueDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// ------
// unions

// SortUnions sorts an array of unions in-place by name.
// This does not sort each union's set of values, use `SortUnionsValues` to sort values.
func SortUnions(nodes []*ast.UnionDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// SortUnionsMap will sort a map of unions by name and return a newly created slice.
// This does not sort each union's set of values, use `SortUnionsValues` to sort values.
func SortUnionsMap(all map[string]*ast.UnionDefinition) []*ast.UnionDefinition {
	var nodes []*ast.UnionDefinition
	for _, one := range all {
		nodes = append(nodes, one)
	}
	SortUnions(nodes)
	return nodes
}

// SortUnionsValues will sort each union's values in-place by name.
func SortUnionsValues(nodes ...*ast.UnionDefinition) {
	for _, one := range nodes {
		SortNamed(one.Types)
	}
}

// -----
// named

// SortNamed sorts an array of named nodes in-place by name.
func SortNamed(nodes []*ast.Named) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// ----------
// interfaces

// SortInterfaces sorts an array of unions in-place by name.
// This does not sort each interface's fields, use `SortInterfacesFields` to sort fields.
func SortInterfaces(nodes []*ast.InterfaceDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// SortInterfacesMap will sort a map of unions by name and return a newly created slice.
// This does not sort each interface's fields, use `SortInterfacesFields` to sort fields.
func SortInterfacesMap(all map[string]*ast.InterfaceDefinition) []*ast.InterfaceDefinition {
	var nodes []*ast.InterfaceDefinition
	for _, one := range all {
		nodes = append(nodes, one)
	}
	SortInterfaces(nodes)
	return nodes
}

// SortInterfacesFields will sort each interface's fields in-place by name.
func SortInterfacesFields(nodes ...*ast.InterfaceDefinition) {
	for _, one := range nodes {
		SortFields(one.Fields)
	}
}

// ----------
// extensions

// SortExtensions sorts an array of type extensions in-place by definition name.
// This does not sort each definition's fields, use `SortExtensionsFields` to sort definition fields.
func SortExtensions(nodes []*ast.TypeExtensionDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Definition.Name.Value,
			nodes[j].Definition.Name.Value,
		})
	})
}

// SortExtensionsMap will sort a map of extensions by name and return a newly created slice.
// This does not sort each interface's fields, use `SortExtensionsFields` to sort fields.
func SortExtensionsMap(all map[string]*ast.TypeExtensionDefinition) []*ast.TypeExtensionDefinition {
	var nodes []*ast.TypeExtensionDefinition
	for _, one := range all {
		nodes = append(nodes, one)
	}
	SortExtensions(nodes)
	return nodes
}

// SortExtensionsFields will sort each extension's definition's fields in-place by name.
func SortExtensionsFields(nodes ...*ast.TypeExtensionDefinition) {
	for _, one := range nodes {
		SortFields(one.Definition.Fields)
	}
}

// ----------
// directives

// SortDirectiveDefinitions sorts an array of directives in-place by name.
// This does not sort each directives' input values, use `SortDirectiveDefinitionsInputValues` to sort input values.
func SortDirectiveDefinitions(nodes []*ast.DirectiveDefinition) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// SortDirectiveDefinitionsMap will sort a map of directive definitions by name and return a newly created slice.
// This does not sort each directives definition's input values, use `SortDirectiveDefinitionsInputValues` to sort input values.
func SortDirectiveDefinitionsMap(all map[string]*ast.DirectiveDefinition) []*ast.DirectiveDefinition {
	var nodes []*ast.DirectiveDefinition
	for _, one := range all {
		nodes = append(nodes, one)
	}
	SortDirectiveDefinitions(nodes)
	return nodes
}

// SortDirectiveDefinitionsInputValues will sort each directive definition's input values in-place by name.
func SortDirectiveDefinitionsInputValues(nodes ...*ast.DirectiveDefinition) {
	for _, one := range nodes {
		SortInputValues(one.Arguments)
	}
}

// SortDirectives sorts an array of directives in-place by name.
// This does not sort each directives' arguments, use `SortDirectivesArguments` to sort arguments.
func SortDirectives(nodes []*ast.Directive) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}

// SortDirectivesMap will sort a map of directives by name and return a newly created slice.
// This does not sort each directive's arguments, use `SortDirectivesArguments` to sort arguments.
func SortDirectivesMap(all map[string]*ast.Directive) []*ast.Directive {
	var nodes []*ast.Directive
	for _, one := range all {
		nodes = append(nodes, one)
	}
	SortDirectives(nodes)
	return nodes
}

// SortDirectivesArguments will sort each directive's arguments in-place by name.
func SortDirectivesArguments(nodes ...*ast.Directive) {
	for _, one := range nodes {
		SortArguments(one.Arguments)
	}
}

// ---------
// arguments

// SortArguments sorts an array of argument nodes in-place by name.
func SortArguments(nodes []*ast.Argument) {
	sort.Slice(nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			nodes[i].Name.Value,
			nodes[j].Name.Value,
		})
	})
}
