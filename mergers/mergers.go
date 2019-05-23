
package mergers

import (
	"fmt"
	"strings"

	"github.com/Fanatics/graphql-ast-helpers/creates"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/richardwilkes/toolbox/errs"
)

// MergeScalars is a simple merge
func MergeScalars(nodes []*ast.ScalarDefinition) (*ast.ScalarDefinition, error) {
	node := &ast.ScalarDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	return node, nil
}

// MergeEnums will apply all the values to eachother
func MergeEnums(nodes []*ast.EnumDefinition) (*ast.EnumDefinition, error) {
	node := &ast.EnumDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	values := make(map[string][]*ast.EnumValueDefinition)
	for _, one := range nodes {
		for _, oneValue := range one.Values {
			curr, _ := values[oneValue.Name.Value]
			values[one.Name.Value] = append(curr, oneValue)
		}
	}

	for _, group := range values {
		value, err := MergeEnumValues(group)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		node.Values = append(node.Values, value)
	}

	return node, nil

}

// MergeEnumValues is a simple merge
func MergeEnumValues(nodes []*ast.EnumValueDefinition) (*ast.EnumValueDefinition, error) {
	node := &ast.EnumValueDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	return node, nil
}

// MergeObjects is a simple merge
func MergeObjects(nodes []*ast.ObjectDefinition) (*ast.ObjectDefinition, error) {
	node := &ast.ObjectDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}
	if node == nil {
		return nil, nil
	}

	var interfaces []string
	fields := make(map[string][]*ast.FieldDefinition)
	for _, one := range nodes {
		for _, ifc := range one.Interfaces {
			interfaces = append(interfaces, ifc.Name.Value)
		}
		for _, field := range one.Fields {
			curr, _ := fields[field.Name.Value]
			fields[field.Name.Value] = append(curr, field)
		}
	}

	for _, one := range uniqueStrings(interfaces) {
		node.Interfaces = append(node.Interfaces,
			creates.NamedType(one),
		)
	}

	for _, group := range fields {
		uniqueFields, err := MergeFields(group)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		node.Fields = append(node.Fields, uniqueFields...)
	}

	return node, nil
}

// MergeInputObjects is a simple merge
func MergeInputObjects(nodes []*ast.InputObjectDefinition) (*ast.InputObjectDefinition, error) {
	node := &ast.InputObjectDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	values := make(map[string][]*ast.InputValueDefinition)
	for _, one := range nodes {
		for _, field := range one.Fields {
			curr, _ := values[field.Name.Value]
			values[field.Name.Value] = append(curr, field)
		}
	}

	for _, group := range values {
		value, err := MergeInputValues(group)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		node.Fields = append(node.Fields, value)
	}

	return node, nil
}

// MergeUnions is a simple merge
func MergeUnions(nodes []*ast.UnionDefinition) (*ast.UnionDefinition, error) {
	node := &ast.UnionDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	var typeNames []string
	for _, one := range nodes {
		for _, named := range one.Types {
			typeNames = append(typeNames, named.Name.Value)
		}
	}
	for _, one := range uniqueStrings(typeNames) {
		node.Types = append(node.Types, creates.NamedType(one))
	}

	return node, nil
}

// MergeInterfaces is a simple merge
func MergeInterfaces(nodes []*ast.InterfaceDefinition) (*ast.InterfaceDefinition, error) {
	node := &ast.InterfaceDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	fields := make(map[string][]*ast.FieldDefinition)
	for _, one := range nodes {
		for _, field := range one.Fields {
			curr, _ := fields[field.Name.Value]
			fields[field.Name.Value] = append(curr, field)
		}
	}

	for _, group := range fields {
		uniqueFields, err := MergeFields(group)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		node.Fields = append(node.Fields, uniqueFields...)
	}

	return node, nil
}

// MergeExtensions is a simple merge
func MergeExtensions(nodes []*ast.TypeExtensionDefinition) (*ast.TypeExtensionDefinition, error) {
	var defs []*ast.ObjectDefinition
	for _, one := range nodes {
		defs = append(defs, one.Definition)
	}

	def, err := MergeObjects(defs)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	if def == nil {
		return nil, nil
	}

	ext := creates.Ext(def.Name.Value)
	ext.Definition = def

	return ext, nil
}

// MergeFields is a simple merge
func MergeFields(nodes []*ast.FieldDefinition) ([]*ast.FieldDefinition, error) {
	all := make(map[string][]*ast.FieldDefinition)
	for _, one := range nodes {
		curr, _ := all[one.Name.Value]
		all[one.Name.Value] = append(curr, one)
	}

	var out []*ast.FieldDefinition

	for _, group := range all {
		node := &ast.FieldDefinition{}
		if err := mergeNodesBasic(group, node); err != nil {
			return nil, errs.Wrap(err)
		}

		types := make(map[string]ast.Type)
		for _, one := range nodes {
			types[fmt.Sprint(printer.Print(one.Type))] = one.Type
		}
		if len(types) > 1 {
			for n, x := range types {
				fmt.Println(n, x)
			}
			return nil, errs.Newf("cannot merge different types, %#v", types)
		}
		for _, nodeType := range types {
			node.Type = nodeType
		}

		args := make(map[string][]*ast.InputValueDefinition)
		for _, one := range nodes {
			for _, arg := range one.Arguments {
				curr, _ := args[one.Name.Value]
				args[one.Name.Value] = append(curr, arg)
			}
		}
		for _, group := range args {
			val, err := MergeInputValues(group)
			if err != nil {
				return nil, errs.Wrap(err)
			}
			node.Arguments = append(node.Arguments, val)
		}

		out = append(out, node)
	}

	return out, nil
}

// MergeDirectivesToOne is a simple merge
func MergeDirectivesToOne(nodes []*ast.DirectiveDefinition) (*ast.DirectiveDefinition, error) {
	node := &ast.DirectiveDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	args := make(map[string][]*ast.InputValueDefinition)
	for _, one := range nodes {
		for _, arg := range one.Arguments {
			curr, _ := args[one.Name.Value]
			args[one.Name.Value] = append(curr, arg)
		}
	}
	for _, group := range args {
		val, err := MergeInputValues(group)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		node.Arguments = append(node.Arguments, val)
	}

	var typeNames []string
	for _, one := range nodes {
		for _, named := range one.Locations {
			typeNames = append(typeNames, named.Value)
		}
	}
	for _, one := range uniqueStrings(typeNames) {
		node.Locations = append(node.Locations, creates.Name(one))
	}

	return node, nil
}

// MergeInputValues is a simple merge
func MergeInputValues(nodes []*ast.InputValueDefinition) (*ast.InputValueDefinition, error) {
	node := &ast.InputValueDefinition{}
	if err := mergeNodesBasic(nodes, node); err != nil {
		return nil, errs.Wrap(err)
	}

	var typeNames []string
	var defaultValues []ast.Value
	for _, one := range nodes {
		typeNames = append(typeNames, one.Type.String())
		defaultValues = append(defaultValues, one.DefaultValue)
	}

	typeName, err := strEqual(typeNames)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	defaultValue, rest := defaultValues[0], defaultValues[1:]
	for len(rest) > 0 {
		step, err := MergeValue(defaultValue, rest[0])
		if err != nil {
			return nil, errs.Wrap(err)
		}
		defaultValue, rest = step, rest[1:]
	}

	node.Type = creates.Type(typeName)
	node.DefaultValue = defaultValue

	return node, nil
}

// MergeNodes attemps to merge multiple nodes of the same type
func MergeNodes(nodes []ast.Node) (ast.Node, error) {
	var allErrs []error

	allKinds := make(map[string]struct{})
	for _, one := range nodes {
		allKinds[one.GetKind()] = struct{}{}
	}

	if len(allKinds) == 0 {
		return nil, nil
	} else if len(allKinds) > 1 {
		return nil, errs.Newf("can only merge nodes of one kind ... %#v", allKinds)
	}

	var out ast.Node

	// just one iteration
	for kind := range allKinds {
		switch kind {
		case kinds.ScalarDefinition:
			var curr []*ast.ScalarDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.ScalarDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeScalars(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		case kinds.EnumDefinition:
			var curr []*ast.EnumDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.EnumDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeEnums(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		case kinds.ObjectDefinition:
			var curr []*ast.ObjectDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.ObjectDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeObjects(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		case kinds.InputObjectDefinition:
			var curr []*ast.InputObjectDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.InputObjectDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeInputObjects(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		case kinds.UnionDefinition:
			var curr []*ast.UnionDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.UnionDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeUnions(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		case kinds.InterfaceDefinition:
			var curr []*ast.InterfaceDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.InterfaceDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeInterfaces(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		case kinds.TypeExtensionDefinition:
			var curr []*ast.TypeExtensionDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.TypeExtensionDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeExtensions(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		case kinds.DirectiveDefinition:
			var curr []*ast.DirectiveDefinition
			for i, one := range nodes {
				if node, ok := one.(*ast.DirectiveDefinition); !ok {
					allErrs = append(allErrs, errs.Newf("node %d was %T, should have been %s", i, node, kind))
					continue
				} else if node != nil {
					curr = append(curr, node)
				}
			}
			if node, err := MergeDirectivesToOne(curr); err != nil {
				allErrs = append(allErrs, errs.Wrap(err))
			} else {
				out = node
			}
		default:
			return nil, errs.Newf("cannot add additional root level type %s", kind)
		}
	}

	return out, nil
}

// MergeLikeDirectives merges directive names and args over multiple definitions
func MergeLikeDirectives(directives []*ast.Directive) ([]*ast.Directive, error) {
	groups := make(map[string][]*ast.Directive)
	for _, one := range directives {
		curr, _ := groups[one.Name.Value]
		groups[one.Name.Value] = append(curr, one)
	}

	var merged []*ast.Directive

	for name, group := range groups {
		one := creates.Directive(name)
		args := make(map[string][]*ast.Argument)
		for _, dir := range group {
			for _, arg := range dir.Arguments {
				curr, _ := args[arg.Name.Value]
				args[arg.Name.Value] = append(curr, arg)
			}
		}
		for argName, argGroup := range args {
			var value ast.Value
			for _, one := range argGroup {
				next, err := MergeValue(value, one.Value)
				if err != nil {
					return nil, errs.Wrap(err)
				}
				value = next
			}
			if value == nil {
				return nil, errs.Newf("@%s(%s ... value not found)", name, argName)
			}
			one.Arguments = append(one.Arguments,
				creates.Arg(argName, value),
			)
		}

		merged = append(merged, one)
	}

	return merged, nil
}

// MergeDirective tries to merge a directive, like-typed values will merge if it makes sense,
// array types will be concatted.
func MergeDirective(directives []*ast.Directive, dir *ast.Directive) ([]*ast.Directive, error) {
	var curr *ast.Directive
	for _, one := range directives {
		if dir.Name.Value == one.Name.Value {
			curr = one
			break
		}
	}

	if curr == nil {
		directives = append(directives, dir)
		return directives, nil
	}

	for _, srcArg := range dir.Arguments {
		foundIdx := -1
		for currIdx, currArg := range curr.Arguments {
			if srcArg.Name.Value == currArg.Name.Value {
				foundIdx = currIdx
				break
			}
		}
		if foundIdx == -1 {
			curr.Arguments = append(curr.Arguments, srcArg)
			continue
		}
		found := curr.Arguments[foundIdx]
		val, err := MergeValue(found.Value, srcArg.Value)
		if err != nil {
			return nil, errs.NewfWithCause(
				err,
				"... on @%q(%s)", dir.Name.Value, found.Name.Value,
			)
		}
		curr.Arguments[foundIdx].Value = val
	}

	return directives, nil
}

// MergeValue tries to merge a value, like-typed values will merge if it makes sense,
// array types will be concatted.
func MergeValue(target, source ast.Value) (ast.Value, error) {
	if source == nil {
		return target, nil
	}
	if target == nil {
		return source, nil
	}

	targetKind := target.GetKind()
	sourceKind := source.GetKind()

	if targetKind != sourceKind {
		return nil, errs.Newf(
			"cannot merge values of types %q and %q\n%s\n%s",
			targetKind, sourceKind, target.GetValue(), source.GetValue(),
		)
	}

	switch tv := target.(type) {
	case *ast.BooleanValue:
		left, right := tv.Value, source.(*ast.BooleanValue).Value
		return creates.ValueBoolean(left && right), nil
	case *ast.EnumValue:
		left, right := tv.Value, source.(*ast.EnumValue).Value
		if left != right {
			return nil, errs.Newf("cannot merge differing enum values: %q and %q", left, right)
		}
		return tv, nil
	case *ast.FloatValue:
		left, right := tv.Value, source.(*ast.FloatValue).Value
		if left != right {
			return nil, errs.Newf("cannot merge differing Float values: %q and %q", left, right)
		}
		return tv, nil
	case *ast.IntValue:
		left, right := tv.Value, source.(*ast.IntValue).Value
		if left != right {
			return nil, errs.Newf("cannot merge differing Int values: %q and %q", left, right)
		}
		return tv, nil
	case *ast.ListValue:
		all := tv.Values
		all = append(all, source.(*ast.ListValue).Values...)
		return creates.ValueList(all...), nil
	case *ast.StringValue:
		left, right := tv.Value, source.(*ast.StringValue).Value
		if left != right {
			return nil, errs.Newf("cannot merge differing String values: %q and %q", left, right)
		}
		return tv, nil
	case *ast.ObjectValue:
		type ofmatch struct {
			left  *ast.ObjectField
			right *ast.ObjectField
		}
		every := make(map[string]ofmatch)
		all := creates.ObjVal()
		for _, one := range tv.Fields {
			every[one.Name.Value] = ofmatch{left: one}
		}
		for _, one := range source.(*ast.ObjectValue).Fields {
			if curr, exists := every[one.Name.Value]; exists {
				curr.right = one
			} else {
				every[one.Name.Value] = ofmatch{right: one}
			}
		}
		for name, match := range every {
			merged, err := MergeValue(match.left.Value, match.right.Value)
			if err != nil {
				return nil, err
			}
			all.Fields = append(all.Fields,
				creates.ObjValField(name, merged),
			)
		}
		return all, nil
	}

	return nil, nil
}



// -------
// helpers

func mergeNodesBasic(raw interface{}, out ast.Node) error {
	var names []string
	var dirs []*ast.Directive
	var descs []*ast.StringValue

	switch nodes := raw.(type) {
	case []*ast.ScalarDefinition:
		vout, ok := out.(*ast.ScalarDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.Scalar(name)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	case []*ast.EnumDefinition:
		vout, ok := out.(*ast.EnumDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.Enum(name, nil)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	case []*ast.EnumValueDefinition:
		vout, ok := out.(*ast.EnumValueDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.EnumVal(name)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	case []*ast.ObjectDefinition:
		vout, ok := out.(*ast.ObjectDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.Obj(name)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	case []*ast.InputObjectDefinition:
		vout, ok := out.(*ast.InputObjectDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.InputObj(name)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	case []*ast.UnionDefinition:
		vout, ok := out.(*ast.UnionDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.Union(name)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	case []*ast.InterfaceDefinition:
		vout, ok := out.(*ast.InterfaceDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.Interface(name)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	// // NOTE!! - just use the internal object definition
	// case []*ast.TypeExtensionDefinition:
	case []*ast.DirectiveDefinition:
		vout, ok := out.(*ast.DirectiveDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.DirectiveDef(name)
		curr.Description = mergeDescriptions(descs)
		*vout = *curr
	case []*ast.FieldDefinition:
		vout, ok := out.(*ast.FieldDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.Field(name, nil)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr
	case []*ast.InputValueDefinition:
		vout, ok := out.(*ast.InputValueDefinition)
		if !ok {
			return errs.Newf("cannot put %T into %T", vout, nodes)
		}
		switch len(nodes) {
		case 0:
			return nil
		case 1:
			*vout = *nodes[0]
			return nil
		}
		for _, one := range nodes {
			names = append(names, one.Name.Value)
			dirs = append(dirs, one.Directives...)
			descs = append(descs, one.Description)
		}
		name, err := strEqual(names)
		if err != nil {
			return errs.Wrap(err)
		}
		directives, err := MergeLikeDirectives(dirs)
		if err != nil {
			return errs.Wrap(err)
		}
		curr := creates.InputVal(name, nil)
		curr.Description = mergeDescriptions(descs)
		curr.Directives = append(curr.Directives, directives...)
		*vout = *curr

	default:
		return errs.Newf("cannot merge type %T", raw)
	}

	return nil
}

func strEqual(strs []string) (string, error) {
	if len(strs) == 0 {
		return "", nil
	}
	first, rest := strs[0], strs[1:]
	for _, one := range rest {
		if first != one {
			return "", errs.Newf("mismatch ... %q and %q", first, one)
		}
	}
	return first, nil
}

func mergeDescriptions(descs []*ast.StringValue) *ast.StringValue {
	if len(descs) == 0 {
		return nil
	}

	var strs []string
	for _, desc := range descs {
		if desc != nil && desc.Value != "" {
			strs = append(strs, desc.Value)
		}
	}

	total := strings.Join(uniqueStrings(strs), "\n\n")

	return creates.ValueString(total)
}

// credit @nilslice for the tip abount empty `struct{}`
func uniqueStrings(in []string) []string {
	var out []string

	all := make(map[string]struct{})
	for _, one := range in {
		if _, has := all[one]; !has {
			out = append(out, one)
			all[one] = struct{}{}
		}
	}

	return out
}