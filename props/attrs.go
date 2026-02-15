package props

import (
	"maps"

	"github.com/a-h/templ"
)

// Attrs wraps templ.Attributes and adds helper functions
// With pattern is used for literal chaining, since the structs are small
type Attrs struct {
	templ.Attributes
}

// With sets given key attribute with string val
func (a Attrs) With(key, val string) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{}
	}
	a.Attributes[key] = val
	return a
}

// Merge merges another Attrs into this one
func (a Attrs) Merge(other Attrs) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{}
	}
	maps.Copy(a.Attributes, other.Attributes)
	return a
}

// WithID sets the id attribute
func (a Attrs) WithID(id string) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{}
	}
	a.Attributes["id"] = id
	return a
}

// WithStyle sets or appends a style attribute
func (a Attrs) WithStyle(style string) Attrs {
    if a.Attributes == nil {
        a.Attributes = templ.Attributes{}
    }

    if existing, ok := a.Attributes["style"]; ok && existing != "" {
        oldStyle, _ := existing.(string)
        // append new style with a semicolon separator
        a.Attributes["style"] = oldStyle + "; " + style
    } else {
        a.Attributes["style"] = style
    }

    return a
}

// WithClass sets or appends a class
func (a Attrs) WithClass(class string) Attrs {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{}
	}
	if existing, ok := a.Attributes["class"]; ok && existing != "" {
		oldClass, _ := existing.(string)
		a.Attributes["class"] = oldClass + " " + class
	} else {
		a.Attributes["class"] = class
	}
	return a
}

// Returns templ attributes for html tags
func (a Attrs) AsTemplAttrs() templ.Attributes {
	if a.Attributes == nil {
		a.Attributes = templ.Attributes{}
	}
	return a.Attributes
}
