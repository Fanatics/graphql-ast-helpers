package bagger

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/richardwilkes/toolbox/errs"
)

// Bagger helpers accumulate a unique (and therefore schema-valid)
// set of each kind of graphql type keyed by node name,
// handily providing helpers to add fields directly to
// top-level query/mutation/subscription types
type Bagger struct {
	scalars map[string]*ast.ScalarDefinition
	enums   map[string]*ast.EnumDefinition
	types   map[string]*ast.ObjectDefinition
	inputs  map[string]*ast.InputObjectDefinition

	unions     map[string]*ast.UnionDefinition
	interfaces map[string]*ast.InterfaceDefinition
	extensions map[string]*ast.TypeExtensionDefinition
	directives map[string]*ast.DirectiveDefinition

	fieldsQuery        map[string]*ast.FieldDefinition
	fieldsMutation     map[string]*ast.FieldDefinition
	fieldsSubscription map[string]*ast.FieldDefinition
}

// NewBagger constructs the bagger and
// allocates the necessary internal maps
func NewBagger() *Bagger {
	b := &Bagger{
		scalars: make(map[string]*ast.ScalarDefinition),
		enums:   make(map[string]*ast.EnumDefinition),
		types:   make(map[string]*ast.ObjectDefinition),
		inputs:  make(map[string]*ast.InputObjectDefinition),

		unions:     make(map[string]*ast.UnionDefinition),
		interfaces: make(map[string]*ast.InterfaceDefinition),
		extensions: make(map[string]*ast.TypeExtensionDefinition),
		directives: make(map[string]*ast.DirectiveDefinition),

		fieldsQuery:        make(map[string]*ast.FieldDefinition),
		fieldsMutation:     make(map[string]*ast.FieldDefinition),
		fieldsSubscription: make(map[string]*ast.FieldDefinition),
	}

	return b
}

// AddNode adds a single node of any kind,
// don't add fields directly here, see one of
// the following:
//
// AddFieldQuery, AddFieldMutation, AddFieldSubscription
func (b *Bagger) AddNode(node ast.Node) error {
	if node == nil {
		return nil
	}

	if err := b.ValidateName(node); err != nil {
		return errs.Wrap(err)
	}

	switch one := node.(type) {
	case *ast.ScalarDefinition:
		name := one.Name.Value
		if _, exists := b.scalars[name]; exists {
			return errs.Newf("scalar %q already exists", name)
		}
		b.scalars[name] = one
	case *ast.EnumDefinition:
		name := one.Name.Value
		if _, exists := b.enums[name]; exists {
			return errs.Newf("enum %q already exists", name)
		}
		b.enums[name] = one
	case *ast.ObjectDefinition:
		name := one.Name.Value
		if _, exists := b.types[name]; exists {
			return errs.Newf("type %q already exists", name)
		}
		b.types[name] = one
	case *ast.InputObjectDefinition:
		name := one.Name.Value
		if _, exists := b.inputs[name]; exists {
			return errs.Newf("input %q already exists", name)
		}
		b.inputs[name] = one
	case *ast.UnionDefinition:
		name := one.Name.Value
		if _, exists := b.unions[name]; exists {
			return errs.Newf("union %q already exists", name)
		}
		b.unions[name] = one
	case *ast.InterfaceDefinition:
		name := one.Name.Value
		if _, exists := b.interfaces[name]; exists {
			return errs.Newf("interface %q already exists", name)
		}
		b.interfaces[name] = one
	case *ast.TypeExtensionDefinition:
		name := one.Definition.Name.Value
		if _, exists := b.extensions[name]; exists {
			return errs.Newf("type extension %q already exists", name)
		}
		b.extensions[name] = one
	case *ast.DirectiveDefinition:
		name := one.Name.Value
		if _, exists := b.directives[name]; exists {
			return errs.Newf("directive %q already exists", name)
		}
		b.directives[name] = one
	case *ast.FieldDefinition:
		return errs.Newf("don't add fields through Bagger.AddNode(), %q", one.Name.Value)
	default:
		return errs.Newf("kind %q not supported by Bagger.AddNode()", one.GetKind())
	}

	return nil
}

// AddFieldQuery adds a single query field definition
func (b *Bagger) AddFieldQuery(field *ast.FieldDefinition) error {
	if err := b.ValidateName(field); err != nil {
		return errs.Wrap(err)
	}
	name := field.Name.Value
	if _, exists := b.fieldsQuery[name]; exists {
		return errs.Newf("query field %q already exists", name)
	}
	b.fieldsQuery[name] = field
	return nil
}

// AddFieldMutation adds a single mutation field definition
func (b *Bagger) AddFieldMutation(field *ast.FieldDefinition) error {
	if err := b.ValidateName(field); err != nil {
		return errs.Wrap(err)
	}
	name := field.Name.Value
	if _, exists := b.fieldsMutation[name]; exists {
		return errs.Newf("query field %q already exists", name)
	}
	b.fieldsMutation[name] = field
	return nil
}

// AddFieldSubscription adds a single subscription field definition
func (b *Bagger) AddFieldSubscription(field *ast.FieldDefinition) error {
	if err := b.ValidateName(field); err != nil {
		return errs.Wrap(err)
	}
	name := field.Name.Value
	if _, exists := b.fieldsSubscription[name]; exists {
		return errs.Newf("query field %q already exists", name)
	}
	b.fieldsSubscription[name] = field
	return nil
}

// GetScalars returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetScalars() map[string]*ast.ScalarDefinition {
	return b.scalars
}

// GetEnums returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetEnums() map[string]*ast.EnumDefinition {
	return b.enums
}

// GetObjects returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetObjects() map[string]*ast.ObjectDefinition {
	return b.types
}

// GetInputObjects returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetInputObjects() map[string]*ast.InputObjectDefinition {
	return b.inputs
}

// GetUnions returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetUnions() map[string]*ast.UnionDefinition {
	return b.unions
}

// GetInterfaces returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetInterfaces() map[string]*ast.InterfaceDefinition {
	return b.interfaces
}

// GetExtensions returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetExtensions() map[string]*ast.TypeExtensionDefinition {
	return b.extensions
}

// GetDirectives returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetDirectives() map[string]*ast.DirectiveDefinition {
	return b.directives
}

// GetFieldsQuery returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetFieldsQuery() map[string]*ast.FieldDefinition {
	return b.fieldsQuery
}

// GetFieldsMutation returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetFieldsMutation() map[string]*ast.FieldDefinition {
	return b.fieldsMutation
}

// GetFieldsSubscription returns the underlying map, which can be directly edited (carefully)
func (b *Bagger) GetFieldsSubscription() map[string]*ast.FieldDefinition {
	return b.fieldsSubscription
}

// ValidateName assures the node's Name is valid
// if a node doesn't have a name, no error is returned
func (b *Bagger) ValidateName(node ast.Node) error {
	var name string

	switch one := node.(type) {
	case *ast.ScalarDefinition:
		name = one.Name.Value
	case *ast.EnumDefinition:
		name = one.Name.Value
	case *ast.ObjectDefinition:
		name = one.Name.Value
	case *ast.InputObjectDefinition:
		name = one.Name.Value
	case *ast.UnionDefinition:
		name = one.Name.Value
	case *ast.InterfaceDefinition:
		name = one.Name.Value
	case *ast.TypeExtensionDefinition:
		name = one.Definition.Name.Value
	default:
		return nil
	}

	if name == "" {
		return errs.Newf("node kind %q cannot have empty name", node.GetKind())
	}

	return nil
}
