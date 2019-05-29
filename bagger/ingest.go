package bagger

import (
	"log"

	"github.com/Fanatics/graphql-ast-helpers/mergers"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/richardwilkes/toolbox/errs"
)

// IngestBags merges all types across 1+ bags into the current bag
func (b *Bagger) IngestBags(bags ...*Bagger) error {
	var allErrs []error
	isErrNil := func(err error) bool {
		if err != nil {
			allErrs = append(allErrs, errs.Wrap(err))
			return false
		}
		return true
	}

	next := NewBagger()
	all := append(bags, b)

	var (
		scalars            = make(map[string][]*ast.ScalarDefinition)
		enums              = make(map[string][]*ast.EnumDefinition)
		objects            = make(map[string][]*ast.ObjectDefinition)
		inputs             = make(map[string][]*ast.InputObjectDefinition)
		unions             = make(map[string][]*ast.UnionDefinition)
		interfaces         = make(map[string][]*ast.InterfaceDefinition)
		extensions         = make(map[string][]*ast.TypeExtensionDefinition)
		directives         = make(map[string][]*ast.DirectiveDefinition)
		fieldsQuery        = make(map[string][]*ast.FieldDefinition)
		fieldsMutation     = make(map[string][]*ast.FieldDefinition)
		fieldsSubscription = make(map[string][]*ast.FieldDefinition)
	)

	// optimize to slurp all at once
	for _, bag := range all {
		for name, node := range bag.GetScalars() {
			curr, _ := scalars[name]
			scalars[name] = append(curr, node)
		}
		for name, node := range bag.GetEnums() {
			curr, _ := enums[name]
			enums[name] = append(curr, node)
		}
		for name, node := range bag.GetObjects() {
			curr, _ := objects[name]
			objects[name] = append(curr, node)
		}
		for name, node := range bag.GetInputObjects() {
			curr, _ := inputs[name]
			inputs[name] = append(curr, node)
		}
		for name, node := range bag.GetUnions() {
			curr, _ := unions[name]
			unions[name] = append(curr, node)
		}
		for name, node := range bag.GetInterfaces() {
			curr, _ := interfaces[name]
			interfaces[name] = append(curr, node)
		}
		for name, node := range bag.GetExtensions() {
			curr, _ := extensions[name]
			extensions[name] = append(curr, node)
		}
		for name, node := range bag.GetDirectives() {
			curr, _ := directives[name]
			directives[name] = append(curr, node)
		}
		for name, node := range bag.GetFieldsQuery() {
			curr, _ := fieldsQuery[name]
			fieldsQuery[name] = append(curr, node)
		}
		for name, node := range bag.GetFieldsMutation() {
			curr, _ := fieldsMutation[name]
			fieldsMutation[name] = append(curr, node)
		}
		for name, node := range bag.GetFieldsSubscription() {
			curr, _ := fieldsSubscription[name]
			fieldsSubscription[name] = append(curr, node)
		}
	}

	// do the merge
	for _, group := range scalars {
		if merged, err := mergers.MergeScalars(group); isErrNil(err) {
			log.Println(group)
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range enums {
		if merged, err := mergers.MergeEnums(group); isErrNil(err) {
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range objects {
		if merged, err := mergers.MergeObjects(group); isErrNil(err) {
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range inputs {
		if merged, err := mergers.MergeInputObjects(group); isErrNil(err) {
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range unions {
		if merged, err := mergers.MergeUnions(group); isErrNil(err) {
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range interfaces {
		if merged, err := mergers.MergeInterfaces(group); isErrNil(err) {
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range extensions {
		if merged, err := mergers.MergeExtensions(group); isErrNil(err) {
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range directives {
		if merged, err := mergers.MergeDirectivesToOne(group); isErrNil(err) {
			isErrNil(next.AddNode(merged))
		}
	}
	for _, group := range fieldsQuery {
		if merged, err := mergers.MergeFieldsToOne(group); isErrNil(err) {
			isErrNil(next.AddFieldQuery(merged))
		}
	}
	for _, group := range fieldsMutation {
		if merged, err := mergers.MergeFieldsToOne(group); isErrNil(err) {
			isErrNil(next.AddFieldMutation(merged))
		}
	}
	for _, group := range fieldsMutation {
		if merged, err := mergers.MergeFieldsToOne(group); isErrNil(err) {
			isErrNil(next.AddFieldSubscription(merged))
		}
	}

	if len(allErrs) > 0 {
		return errs.Append(nil, allErrs...)
	}

	*b = *next
	return nil
}
