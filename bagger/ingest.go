package bagger

import (

	// "github.com/Fanatics/graphql-ast-helpers/sorters"
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

	// optimize to do all at once
	for _, bag := range append(bags, b) {

		// scalars
		scalars := make(map[string][]*ast.ScalarDefinition)
		for name, node := range bag.GetScalars() {
			curr, _ := scalars[name]
			scalars[name] = append(curr, node)
		}
		for _, group := range scalars {
			if merged, err := mergers.MergeScalars(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// enums
		enums := make(map[string][]*ast.EnumDefinition)
		for name, node := range bag.GetEnums() {
			curr, _ := enums[name]
			enums[name] = append(curr, node)
		}
		for _, group := range enums {
			if merged, err := mergers.MergeEnums(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// objects
		objects := make(map[string][]*ast.ObjectDefinition)
		for name, node := range bag.GetObjects() {
			curr, _ := objects[name]
			objects[name] = append(curr, node)
		}
		for _, group := range objects {
			if merged, err := mergers.MergeObjects(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// inputs
		inputs := make(map[string][]*ast.InputObjectDefinition)
		for name, node := range bag.GetInputObjects() {
			curr, _ := inputs[name]
			inputs[name] = append(curr, node)
		}
		for _, group := range objects {
			if merged, err := mergers.MergeObjects(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// unions
		unions := make(map[string][]*ast.UnionDefinition)
		for name, node := range bag.GetUnions() {
			curr, _ := unions[name]
			unions[name] = append(curr, node)
		}
		for _, group := range unions {
			if merged, err := mergers.MergeUnions(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// interfaces
		interfaces := make(map[string][]*ast.InterfaceDefinition)
		for name, node := range bag.GetInterfaces() {
			curr, _ := interfaces[name]
			interfaces[name] = append(curr, node)
		}
		for _, group := range interfaces {
			if merged, err := mergers.MergeInterfaces(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// extensions
		extensions := make(map[string][]*ast.TypeExtensionDefinition)
		for name, node := range bag.GetExtensions() {
			curr, _ := extensions[name]
			extensions[name] = append(curr, node)
		}
		for _, group := range extensions {
			if merged, err := mergers.MergeExtensions(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// directives
		directives := make(map[string][]*ast.DirectiveDefinition)
		for name, node := range bag.GetDirectives() {
			curr, _ := directives[name]
			directives[name] = append(curr, node)
		}
		for _, group := range directives {
			if merged, err := mergers.MergeDirectivesToOne(group); !isErrNil(err) {
				isErrNil(next.AddNode(merged))
			}
		}

		// fields - query
		fieldsQuery := make(map[string][]*ast.FieldDefinition)
		for name, node := range bag.GetFieldsQuery() {
			curr, _ := fieldsQuery[name]
			fieldsQuery[name] = append(curr, node)
		}
		for _, group := range fieldsQuery {
			if merged, err := mergers.MergeFieldsToOne(group); !isErrNil(err) {
				isErrNil(next.AddFieldQuery(merged))
			}
		}

		// fields - mutation
		fieldsMutation := make(map[string][]*ast.FieldDefinition)
		for name, node := range bag.GetFieldsQuery() {
			curr, _ := fieldsMutation[name]
			fieldsMutation[name] = append(curr, node)
		}
		for _, group := range fieldsQuery {
			if merged, err := mergers.MergeFieldsToOne(group); !isErrNil(err) {
				isErrNil(next.AddFieldMutation(merged))
			}
		}

		// fields - subscription
		fieldsSubscription := make(map[string][]*ast.FieldDefinition)
		for name, node := range bag.GetFieldsQuery() {
			curr, _ := fieldsSubscription[name]
			fieldsSubscription[name] = append(curr, node)
		}
		for _, group := range fieldsQuery {
			if merged, err := mergers.MergeFieldsToOne(group); !isErrNil(err) {
				isErrNil(next.AddFieldSubscription(merged))
			}
		}
	}

	if len(allErrs) > 0 {
		return errs.Append(nil, allErrs...)
	}

	*b = *next
	return nil
}
