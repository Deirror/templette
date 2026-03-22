// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package props

import (
	"maps"

	"github.com/a-h/templ"
)

// Attrs wraps templ.Attributes and adds helper functions.
type Attrs struct {
	templ.Attributes
}

// With sets given key attribute with string val.
func (a Attrs) With(key, val string) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{} // initialize nil map
	}

	a.Attributes[key] = val
	return a
}

// Merge merges another Attrs into this one.
func (a Attrs) Merge(other Attrs) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{} // initialize nil map
	}

	maps.Copy(a.Attributes, other.Attributes)
	return a
}

// WithID sets the id attribute.
func (a Attrs) WithID(id string) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{} // initialize nil map
	}

	a.Attributes[ID] = id
	return a
}

// WithStyle sets or appends a style attribute.
func (a Attrs) WithStyle(style string) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{} // initialize nil map
	}

	if existing, ok := a.Attributes[Style]; ok && existing != "" {
		oldStyle, _ := existing.(string)
		// append new style with a semicolon separator
		a.Attributes[Style] = oldStyle + "; " + style
	} else {
		a.Attributes[Style] = style // fist style
	}

	return a
}

// WithClass sets or appends a class.
func (a Attrs) WithClass(class string) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{} // initialize nil map
	}

	if existing, ok := a.Attributes[Class]; ok && existing != "" {
		oldClass, _ := existing.(string)
		// append new style with a space separator
		a.Attributes[Class] = oldClass + " " + class
	} else {
		a.Attributes[Class] = class // first class
	}

	return a
}

// AsTemplAttrs flattens all nested Props (Aria, Hx, Data, Attrs) into a single set of templ.Attributes.
func (a Attrs) AsTemplAttrs() templ.Attributes {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{} // initialize nil map
	}

	return a.Attributes
}
