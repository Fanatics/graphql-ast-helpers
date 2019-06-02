// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/richardwilkes/toolbox/errs"
)

var genKinds = map[string]reflect.Type{
	// Name
	kinds.Name: reflect.TypeOf(ast.NewName(nil)),

	// Document
	// kinds.Document:            reflect.TypeOf(ast.NewDocument(nil)),
	// kinds.OperationDefinition: reflect.TypeOf(ast.NewOperationDefinition(nil)),
	// kinds.VariableDefinition:  reflect.TypeOf(ast.NewVariableDefinition(nil)),
	// kinds.Variable:            reflect.TypeOf(ast.NewVariable(nil)),
	// kinds.SelectionSet:        reflect.TypeOf(ast.NewSelectionSet(nil)),
	kinds.Field:    reflect.TypeOf(ast.NewField(nil)),
	kinds.Argument: reflect.TypeOf(ast.NewArgument(nil)),

	// Fragments
	// kinds.FragmentSpread:     reflect.TypeOf(ast.NewFragmentSpread(nil)),
	// kinds.InlineFragment:     reflect.TypeOf(ast.NewInlineFragment(nil)),
	// kinds.FragmentDefinition: reflect.TypeOf(ast.NewFragmentDefinition(nil)),

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
	// kinds.SchemaDefinition:        reflect.TypeOf(ast.NewSchemaDefinition(nil)),
	// kinds.OperationTypeDefinition: reflect.TypeOf(ast.NewOperationTypeDefinition(nil)),

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

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Funcs(funcs).Parse(tmplstr))
}

func main() {
	mergersDir := os.Args[1]

	for kind, t := range genKinds {
		file, err := os.Create(mergersDir + "/mergers_" + strings.ToLower(kind) + "_gen.go")
		if err != nil {
			log.Fatal(1, err)
		}
		defer file.Close()

		err = tmpl.Execute(file, genKind{
			Kind: kind,
			Type: t,
		})
		if err != nil {
			log.Fatal(1, err)
		}
	}
}

// -------------
// template functions

type genKind struct {
	Kind string
	Type reflect.Type
}

func (g genKind) IsStructy() bool {
	t := g.Type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Kind() == reflect.Struct 
}

func (g genKind) eachField(fn func(reflect.StructField)) error {
	if !g.IsStructy() {
		return errs.Newf("can't each-field a struct")
	}

	elem := g.Type.Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		switch field.Name {
		case "Kind", "Loc":
			continue
		default:
			fn(elem.Field(i))
		}
	}

	return nil
}

func isMulti(sf reflect.StructField) bool {
	sk := sf.Type.Kind()
	return sk == reflect.Slice || sk == reflect.Array
}

func typeAsMulti(sf reflect.StructField) reflect.Type {
	if !isMulti(sf) {
		return reflect.SliceOf(sf.Type)
	}
	return sf.Type
}

var funcs = template.FuncMap{
	"getType":         getType,
	"accessName":      accessName,
	"propsSlices":     propsSlices,
	"propsAppenders":        propsAppenders,
	"propsMerge": propsMerge,
	// "propDirective":   propDirective,
	// "propInterface":   propInterface,
}

var propsSlices = func(one genKind) (string, error) {
	var decls []string

	one.eachField(func(field reflect.StructField) {
		ft := typeAsMulti(field)

		decls = append(decls,
			fmt.Sprintf("  var list%s %s", field.Name, ft),
		)
	})

	return strings.Join(decls, "\n"), nil
}

var propsAppenders = func(one genKind) (string, error) {
	var appenders []string

	one.eachField(func(field reflect.StructField) {
		suffix := ""
		if isMulti(field) {
			suffix = "..."
		}

		fname := field.Name
		appenders = append(appenders,
			fmt.Sprintf("    list%s = append(list%s, one.%s%s)", fname, fname, fname, suffix),
		)
	})

	return strings.Join(appenders, "\n"), nil
}

var tmplMerge = template.Must(template.New("").Parse(`
	{{- /*etc */}}  if merged, err := m.{{ .Fn }}(list{{ .Fname }}); err != nil {
		errSet = errs.Append(errSet, err)
	} else {
		one.{{ .Fname }} = merged
	}`))

var propsMerge = func(one genKind) (string, error) {
	var appenders []string
	var errSet error

	one.eachField(func(field reflect.StructField) {
		out := &strings.Builder{}

		pcs := strings.Split(field.Type.String(), ".")

		fn := pcs[len(pcs) - 1]
		if isMulti(field) {
			fn = "Similar" + fn
		} else {
			fn = "One" + fn
		}

		if err := tmplMerge.Execute(out, struct{
			Fn string
			Fname string
		}{
			Fn: fn,
			Fname: field.Name,
		}); err != nil {
			errSet = errs.Append(errSet, err)
		}

		appenders = append(appenders, out.String())
	})

	return strings.Join(appenders, "\n"), errSet
}

var getConstructor = func(one genKind) (string, error) {
	if one.IsStructy() {
		return fmt.Sprintf("ast.New%s(nil)", one.Kind), nil
	}

	return "", nil
}

var getType = func(one genKind) string {
	return one.Type.String()
}

var accessName = func(varname, src string, one genKind) string {
	log.Println(one.Type.Name(), one.Type.String())
	switch one.Type.String() {
	case "ast.Value", "*ast.EnumValue", "*ast.NonNull", "*ast.List", "*ast.ObjectValue", "*ast.FloatValue", "*ast.IntValue", "*ast.StringValue":
		return fmt.Sprintf("%s := m.getValueID(%s)", varname, src)
	case "ast.Type":
		return fmt.Sprintf("%s := fmt.Sprint(printer.Print(%s))", varname, src)
	}
	// if one.Type.Kind() == reflect.Ptr {
	// 	if one.Type.Implements(reflect.TypeOf(new(ast.Value)).Elem()) {
	// 		return fmt.Sprintf("%s := m.getValueID(%s)", varname, src)
	// 	}
	// 	if one.Type.Implements(reflect.TypeOf(new(ast.Type)).Elem()) {
	// 		return fmt.Sprintf("%s := fmt.Sprint(printer.Print(%s))", varname, src)
	// 	}
	// }

	if name := propName(one, src); name != "" {
		return fmt.Sprintf("%s := fmt.Sprint(printer.Print(%s))", varname, name)
	}

	switch one.Kind {
	case kinds.Name:
		return fmt.Sprintf("%s := fmt.Sprint(printer.Print(%s))", varname, src)
	case kinds.TypeExtensionDefinition:
		return fmt.Sprintf("%s := fmt.Sprint(printer.Print(%s.Definition.Name))", varname, src)
	}

	log.Fatalf("kind %q not supported", one.Kind)
	return ""
}

// // var constructKind = func(one genKind) string {
// // }

// // property accessors

var propName = func(one genKind, varname string) string {
	switch one.Kind {
	case kinds.Name:
		return varname
	case kinds.TypeExtensionDefinition:
		return fmt.Sprintf("%s.Definition.Name", varname)
	}

	if hasFieldKind(one.Type, genKinds[kinds.Name], "Name") {
		return fmt.Sprintf("%s.Name", varname)
	}

	return ""
}

// var propDescription = func(one genKind, varname string) string {
// 	switch one.Kind {
// 	case kinds.TypeExtensionDefinition:
// 		return fmt.Sprintf("%s.Definition.Description", varname)
// 	}

// 	if hasFieldKind(one.Value, genKinds[kinds.StringValue].Type(), "Description") {
// 		return fmt.Sprintf("%s.Description", varname)
// 	}

// 	return ""
// }

// var sliceDirectives = reflect.SliceOf(genKinds[kinds.Directive].Type())

// var propDirective = func(one genKind, varname string) string {
// 	switch one.Kind {
// 	case kinds.TypeExtensionDefinition:
// 		return fmt.Sprintf("%s.Definition.Directives", varname)
// 	}

// 	if hasFieldKind(one.Value, sliceDirectives, "Directives") {
// 		return fmt.Sprintf("%s.Directives", varname)
// 	}

// 	return ""
// }

// var propInterface = func(one genKind, varname string) string {
// 	switch one.Kind {
// 	case kinds.ObjectDefinition:
// 		return fmt.Sprintf("%s.Interfaces", varname)
// 	case kinds.TypeExtensionDefinition:
// 		return fmt.Sprintf("%s.Definition.Interfaces", varname)
// 	}
// 	return ""
// }

// -------
// helpers

func hasFieldKind(parent, child reflect.Type, name string) bool {
	if parent.Kind() == reflect.Ptr {
		if field, found := parent.Elem().FieldByName(name); found {
			return field.Type == child
		}
	}
	return false
}

// ---------------
// template string

// TODO: needs more
// {{- $propValue, $readValue := propValue . }}
// {{- $propFieldsObject, $readFieldsObject := propFieldsObject . }}
// {{- $propFieldsInputObject, $readFieldsInputObject := propFieldsInputObject . }}

var tmplstr = `// Code generated by go generate; DO NOT EDIT.
package mergers

import (
	"fmt"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
	"github.com/richardwilkes/toolbox/errs"
)

var _ = fmt.Sprint
var _ = printer.Print

{{- $type := getType .}}

// Similar{{ .Kind }} merges declarations of {{ .Kind }} that share the same {{ .Kind }} value.
func (m *Merger) Similar{{ .Kind }}(curr []{{ $type }}, more ...{{ $type }}) ([]{{ $type }}, error) {
	all := append(curr, more...)
	if len(all) <= 1 {
		return all, nil
	}

	groups := make(map[string][]{{ $type }})
	for _, one := range all {
		{{ accessName "name" "one" . }}
		if name != "" {
			curr, _ := groups[name]
			groups[name] = append(curr, one)
		}
	}

	var out []{{ $type }}
	var errSet error

	for _, group := range groups {
		if merged, err := m.One{{ .Kind }}(group); err != nil {
			errSet = errs.Append(errSet, err)
		} else if merged != nil {
			out = append(out, merged)
		}
	}

	return out, errSet
}

// One{{ .Kind }} attempts to merge all members of {{ .Kind }} into a singe {{ $type }}.
// If this cannot be done, this method will return an error.
func (m *Merger) One{{ .Kind }}(curr []{{ $type }}, more ...{{ $type }}) ({{ $type }}, error) {
	// step 1 - escape hatch when no calculation is needed
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return nil, nil
	} else if n == 1 {
		return all[0], nil
	}

	// step 2 - prepare property collections (if any)
{{ propsSlices . }}

	// step 3 - range over the parent struct and collect properties
	{{- $pm := propsAppenders . }}
	{{- if $pm }}
	for _, one := range all {
{{ $pm }}
	}
	{{- end }}

	// step 4 - prepare output types
	one := ast.New{{ .Kind }}(nil)
	var errSet error

	// step 5 - merge properties
{{ propsMerge . }}

	return one, errSet
}
`
