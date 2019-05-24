package bagger

import (
	"github.com/Fanatics/graphql-ast-helpers/creates"
	"github.com/Fanatics/graphql-ast-helpers/sorters"
	"github.com/graphql-go/graphql/language/ast"
)

// Export writes the contents of the Bager to a document,
// sorting everything along the way
func (b *Bagger) Export(
	queryObject, mutationObject, subscriptionObject *ast.ObjectDefinition,
) *ast.Document {
	doc := ast.NewDocument(nil)
	addNode := func(node ast.Node) {
		doc.Definitions = append(doc.Definitions, node)
	}

	var (
		queryName        = maybeGetObjectName(queryObject)
		mutationName     = maybeGetObjectName(mutationObject)
		subscriptionName = maybeGetObjectName(subscriptionObject)
	)

	// schema
	addNode(creates.Schema(queryName, mutationName, subscriptionName))

	// top-level objects
	if queryObject != nil {
		for _, one := range sorters.SortFieldsMap(b.GetFieldsQuery()) {
			queryObject.Fields = append(queryObject.Fields, one)
		}
		addNode(queryObject)
	}
	if mutationObject != nil {
		for _, one := range sorters.SortFieldsMap(b.GetFieldsMutation()) {
			mutationObject.Fields = append(mutationObject.Fields, one)
		}
		addNode(mutationObject)
	}
	if subscriptionObject != nil {
		for _, one := range sorters.SortFieldsMap(b.GetFieldsSubscription()) {
			subscriptionObject.Fields = append(subscriptionObject.Fields, one)
		}
		addNode(subscriptionObject)
	}

	// type definitions
	for _, one := range sorters.SortScalarsMap(b.GetScalars()) {
		addNode(one)
	}
	for _, one := range sorters.SortEnumsMap(b.GetEnums()) {
		addNode(one)
	}
	for _, one := range sorters.SortObjectsMap(b.GetObjects()) {
		sorters.SortObjectsFields(one)
		addNode(one)
	}
	for _, one := range sorters.SortInputObjectsMap(b.GetInputObjects()) {
		sorters.SortInputObjectsFields(one)
		addNode(one)
	}
	for _, one := range sorters.SortUnionsMap(b.GetUnions()) {
		sorters.SortUnionsValues(one)
		addNode(one)
	}
	for _, one := range sorters.SortInterfacesMap(b.GetInterfaces()) {
		sorters.SortInterfacesFields(one)
		addNode(one)
	}
	for _, one := range sorters.SortExtensionsMap(b.GetExtensions()) {
		sorters.SortExtensionsFields(one)
		addNode(one)
	}
	for _, one := range sorters.SortDirectiveDefinitionsMap(b.GetDirectives()) {
		sorters.SortDirectiveDefinitionsInputValues(one)
		addNode(one)
	}

	return doc
}

func maybeGetObjectName(obj *ast.ObjectDefinition) string {
	if obj == nil || obj.Name == nil {
		return ""
	}
	return obj.Name.Value
}
