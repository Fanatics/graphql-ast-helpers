package mergers

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
	"github.com/richardwilkes/toolbox/errs"
)

// Basic is the default merger strategy
var Basic = &Merger{
	getValueID: defaultValueID,
	getNodeID:  defaultNodeID,
}

// Merger ...
type Merger struct {
	getValueID func(val ast.Value) string
	getNodeID  func(node ast.Node) string
}

// NewMerger ...
func NewMerger(opts ...func(*Merger) error) (*Merger, error) {
	m := &Merger{
		getValueID: defaultValueID,
		getNodeID:  defaultNodeID,
	}

	for _, one := range opts {
		if err := one(m); err != nil {
			return nil, errs.Wrap(err)
		}
	}

	return m, nil
}

// Onestring merges strings
func (m *Merger) Onestring(curr []string, more ...string) (string, error) {
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return "", nil
	} else if n == 1 {
		return all[0], nil
	}

	group := make(map[string]struct{})
	for _, one := range all {
		group[one] = struct{}{}
	}

	if len(group) > 1 {
		return "", errs.Newf("multiple different strings %#v", group)
	}

	var out string
	for str := range group {
		out = str
	}
	return out, nil
}

// Onebool merges strings
func (m *Merger) Onebool(curr []bool, more ...bool) (bool, error) {
	all := append(curr, more...)
	if n := len(all); n == 0 {
		return false, nil
	} else if n == 1 {
		return all[0], nil
	}

	group := make(map[bool]struct{})
	for _, one := range all {
		group[one] = struct{}{}
	}

	if len(group) > 1 {
		return false, errs.Newf("multiple different bools %#v", group)
	}

	var out bool
	for b := range group {
		out = b
	}
	return out, nil
}

// ---------------------
// helpers and constants

var defaultValueID = func(val ast.Value) string {
	return fmt.Sprint(printer.Print(val))
}

var defaultNodeID = func(node ast.Node) string {
	return fmt.Sprint(printer.Print(node))
}
