package props

import "github.com/a-h/templ"

// Attrs wraps templ.Attributes and adds helper functions
type Attrs struct {
	templ.Attributes
}

// NewAttrs creates a new Attrs
func NewAttrs() Attrs {
	return Attrs{Attributes: templ.Attributes{}}
}

// With sets given key attribute with string val
func (a Attrs) With(key, val string) Attrs {
	a.Attributes[key] = val
	return a
}

// WithBool sets given key attribute with bool val
func (a Attrs) WithBool(key string, val bool) Attrs {
	a.Attributes[key] = val
	return a
}

// WithID sets the id attribute
func (a Attrs) WithID(id string) Attrs {
	a.Attributes["id"] = id
	return a
}

// WithClass sets or appends a class
func (a Attrs) WithClass(class string) Attrs {
	if existing, ok := a.Attributes["class"]; ok && existing != "" {
		oldClass, _ := existing.(string)
		a.Attributes["class"] = oldClass + " " + class
	} else {
		a.Attributes["class"] = class
	}
	return a
}

// Merge merges another Attrs into this one
func (a Attrs) Merge(other Attrs) Attrs {
	for k, v := range other.Attributes {
		a.Attributes[k] = v
	}
	return a
}
