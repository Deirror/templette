// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package body

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

// Props wraps props.Attrs to provide a typed container for <body> attributes.
type Props struct {
	props.Attrs
}

// AsTemplAttrs returns the wrapped attributes as templ.Attributes for use
// with templ elements.
func (p Props) AsTemplAttrs() templ.Attributes {
	return p.Attrs.AsTemplAttrs()
}
