package generator

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// A Name describes an identifier of an Entity (Message, Field, Enum, Service,
// Field). It can be converted to multiple forms using the provided helper
// methods, or a custom transform can be used to modify its behavior.
type Name string

// String satisfies the strings.Stringer interface.
func (n Name) String() string { return string(n) }

// UpperCamelCase converts Name n to upper camelcase, where each part is
// title-cased and concatenated with no separator.
func (n Name) UpperCamelCase() Name { return n.Transform(strings.Title, strings.Title, "") }

// LowerCamelCase converts Name n to lower camelcase, where each part is
// title-cased and concatenated with no separator except the first which is
// lower-cased.
func (n Name) LowerCamelCase() Name { return n.Transform(strings.Title, strings.ToLower, "") }

// ScreamingSnakeCase converts Name n to screaming-snake-case, where each part
// is all-caps and concatenated with underscores.
func (n Name) ScreamingSnakeCase() Name { return n.Transform(strings.ToUpper, strings.ToUpper, "_") }

// LowerSnakeCase converts Name n to lower-snake-case, where each part is
// lower-cased and concatenated with underscores.
func (n Name) LowerSnakeCase() Name { return n.Transform(strings.ToLower, strings.ToLower, "_") }

// UpperSnakeCase converts Name n to upper-snake-case, where each part is
// title-cased and concatenated with underscores.
func (n Name) UpperSnakeCase() Name { return n.Transform(strings.Title, strings.Title, "_") }

// SnakeCase converts Name n to snake-case, where each part preserves its
// capitalization and concatenated with underscores.
func (n Name) SnakeCase() Name { return n.Transform(ID, ID, "_") }

// LowerDotNotation converts Name n to lower dot notation, where each part is
// lower-cased and concatenated with periods.
func (n Name) LowerDotNotation() Name { return n.Transform(strings.ToLower, strings.ToLower, ".") }

// UpperDotNotation converts Name n to upper dot notation, where each part is
// title-cased and concatenated with periods.
func (n Name) UpperDotNotation() Name { return n.Transform(strings.Title, strings.Title, ".") }

// Split breaks apart Name n into its constituent components. Precedence
// follows dot notation, then underscores (excluding underscore prefixes), then
// camelcase. Numbers are treated as standalone components.
func (n Name) Split() (parts []string) {
	ns := string(n)

	switch {
	case ns == "":
		return []string{""}
	case strings.LastIndex(ns, ".") >= 0:
		return strings.Split(ns, ".")
	case strings.LastIndex(ns, "_") > 0: // leading underscore does not count
		parts = strings.Split(ns, "_")
		if parts[0] == "" {
			parts[1] = "_" + parts[1]
			return parts[1:]
		}
		return
	default: // camelCase
		buf := &bytes.Buffer{}
		var capt, lodash, num bool
		for _, r := range ns {
			uc := unicode.IsUpper(r) || unicode.IsTitle(r)
			dg := unicode.IsDigit(r)

			if r == '_' && buf.Len() == 0 && len(parts) == 0 {
				lodash = true
			}

			if uc && !capt && buf.Len() > 0 && !lodash { // new upper letter
				parts = append(parts, buf.String())
				buf.Reset()
			} else if dg && !num && buf.Len() > 0 && !lodash { // new digit
				parts = append(parts, buf.String())
				buf.Reset()
			} else if !uc && capt && buf.Len() > 1 { // upper to lower
				if ss := buf.String(); len(ss) > 1 &&
					(len(ss) != 2 || ss[0] != '_') {
					pr, _ := utf8.DecodeLastRuneInString(ss)
					parts = append(parts, strings.TrimSuffix(ss, string(pr)))
					buf.Reset()
					buf.WriteRune(pr)
				}
			} else if !dg && num && buf.Len() >= 1 {
				parts = append(parts, buf.String())
				buf.Reset()
			}

			num = dg
			capt = uc
			buf.WriteRune(r)
		}
		parts = append(parts, buf.String())
		return
	}
}

// NameTransformer is a function that mutates a string. Many of the methods in
// the standard strings package satisfy this signature.
type NameTransformer func(string) string

// ID is a NameTransformer that does not mutate the string.
func ID(s string) string { return s }

// Chain combines the behavior of two Transformers into one. If multiple
// transformations need to be performed on a Name, this method should be used
// to reduce it to a single transformation before applying.
func (n NameTransformer) Chain(t NameTransformer) NameTransformer {
	return func(s string) string { return t(n(s)) }
}

// Transform applies a transformation to the parts of Name n, returning a new
// Name. Transformer first is applied to the first part, with mod applied to
// all subsequent ones. The parts are then concatenated with the separator sep.
// For optimal efficiency, multiple NameTransformers should be Chained together
// before calling Transform.
func (n Name) Transform(mod, first NameTransformer, sep string) Name {
	parts := n.Split()

	for i, p := range parts {
		if i == 0 {
			parts[i] = first(p)
		} else {
			parts[i] = mod(p)
		}
	}

	return Name(strings.Join(parts, sep))
}
