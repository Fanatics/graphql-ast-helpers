package template

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/Fanatics/graphql-ast-helpers/meta"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/richardwilkes/toolbox/errs"
)

// MergersFuncs is a funcmap for generated merger code
var MergersFuncs = template.FuncMap{
	"getKey": func(kind, varname string) (string, error) {
		// 1 - interface types
		if isIfc, _, err := meta.IsInterface(kind); err != nil {
			return "", errs.Wrap(err)
		} else if isIfc {
			switch kind {
			case "Value":
				return fmt.Sprintf("m.getValueID(%s)", varname), nil
			case "Type":
				return fmt.Sprintf("fmt.Sprint(printer.Print(%s))", varname), nil
			case "Selection":
				return `"selection"`, nil
			case "Node":
				return fmt.Sprintf("m.getNodeID(%s)", varname), nil
			}
			return "", errs.Newf("interface kind %q is not known", kind)
		}

		// 2 - specific cases of concrete types
		switch kind {
		case kinds.Name:
			return fmt.Sprintf("fmt.Sprint(printer.Print(%s))", varname), nil
		case kinds.TypeExtensionDefinition:
			return fmt.Sprintf("fmt.Sprint(printer.Print(%s.Definition.Name))", varname), nil
		case kinds.Document:
			return fmt.Sprintf(`"document"`), nil
		}

		// 3 - concrete types with a "name" field
		if hasName, _, err := meta.HasFieldKind(kind, kinds.Name, "Name"); err != nil {
			return "", errs.Wrap(err)
		} else if hasName {
			return fmt.Sprintf("fmt.Sprint(printer.Print(%s.Name))", varname), nil
		}

		// 4 - any remaining node type
		if isNode, err := meta.DoesImplement(kind, "Node"); err != nil {
			return "", errs.Wrap(err)
		} else if isNode {
			return fmt.Sprintf("m.getNodeID(%s)", varname), nil
		}

		return "", errs.Newf("kind %q is not known", kind)
	},

	"propsSlices": func(kind string) (string, error) {
		if isInterface, _, err := meta.IsInterface(kind); err != nil {
			return "", errs.Wrap(err)
		} else if isInterface {
			return "", nil
		}

		var decls []string

		fields, err := meta.AllFields(kind)
		if err != nil {
			return "", errs.Wrap(err)
		}

		for _, field := range fields {
			ft := typeAsMulti(field)
			decl := fmt.Sprintf("var list%s %s", field.Name, ft)
			decls = append(decls, decl)
		}

		return strings.Join(decls, "\n"), nil
	},

	"propsAppenders": func(kind, varname string) (string, error) {
		if isInterface, _, err := meta.IsInterface(kind); err != nil {
			return "", errs.Wrap(err)
		} else if isInterface {
			return "", nil
		}

		var appenders []string

		fields, err := meta.AllFields(kind)
		if err != nil {
			return "", errs.Wrap(err)
		}

		for _, field := range fields {
			suffix := ""
			if isMulti(field.Type) {
				suffix = "..."
			}

			fname := field.Name
			appender := fmt.Sprintf("list%s = append(list%s, %s.%s%s)", fname, fname, varname, fname, suffix)
			appenders = append(appenders, appender)
		}

		return strings.Join(appenders, "\n"), nil
	},

	"propsMergers": func(kind string) (string, error) {
		out := &strings.Builder{}

		data := tmplMerge{
			Kind: kind,
		}

		isConcrete, _, err := meta.IsConcrete(kind)
		if err != nil {
			return "", errs.Wrap(err)
		}
		if isConcrete {
			fields, err := meta.AllFields(kind)
			if err != nil {
				return "", errs.Wrap(err)
			}
			data.Fields = fields
			if err := tmplConcreteMerge.Execute(out, data); err != nil {
				return "", errs.Wrap(err)
			}
			return out.String(), nil
		}

		isInterface, _, err := meta.IsInterface(kind)
		if err != nil {
			return "", errs.Wrap(err)
		}
		if isInterface {
			implementers, err := meta.AllImplementers(kind)
			if err != nil {
				return "", errs.Wrap(err)
			}
			data.Implementers = implementers
			if err := tmplIfcMerge.Execute(out, data); err != nil {
				return "", errs.Wrap(err)
			}
			return out.String(), nil
		}

		return "", errs.Newf("kind %q not concrete or interface", kind)
	},
}

type tmplMerge struct {
	Kind         string
	Fields       []reflect.StructField
	Implementers map[string]reflect.Type
}

var tmplMergeFuncMap = template.FuncMap{
	"getFn": func(rtype reflect.Type) (string, error) {
		pcs := strings.Split(rtype.String(), ".")
		fn := pcs[len(pcs)-1]
		if isMulti(rtype) {
			fn = "Similar" + fn
		} else {
			fn = "One" + fn
		}
		return fn, nil
	},
}

var tmplConcreteMerge = template.Must(template.New("").Funcs(tmplMergeFuncMap).Parse(`
one := ast.New{{ .Kind }}(nil)

{{- range .Fields }}
if merged, err := m.{{ getFn .Type }}(list{{ .Name }}); err != nil {
	errSet = errs.Append(errSet, err)
} else {
	one.{{ .Name }} = merged
}
{{- end }}

return one, errSet
`))

var tmplIfcMerge = template.Must(template.New("").Funcs(tmplMergeFuncMap).Parse(`
switch all[0].(type) {
{{- range $k, $v := .Implementers }}
case {{ $v }}:
	var set []{{ $v }}
	for _, single := range all {
		v, ok := single.({{ $v }})
		if !ok {
			errSet = errs.Append(errSet, errs.Newf("want {{ $v }} but got type %T", single))
			continue
		}
		set = append(set, v)
	}
	return m.{{ getFn $v }}(set)
{{- end }}
default:
	errSet = errs.Append(errSet, errs.Newf("type %T unknown", all[0]))
}

return nil, errSet
`))

// -------
// helpers
func isMulti(rtype reflect.Type) bool {
	rk := rtype.Kind()
	return rk == reflect.Slice || rk == reflect.Array
}

func typeAsMulti(sf reflect.StructField) reflect.Type {
	if !isMulti(sf.Type) {
		return reflect.SliceOf(sf.Type)
	}
	return sf.Type
}
