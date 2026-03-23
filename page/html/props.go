// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package html

import (
	"github.com/Deirror/templette/props"
	"github.com/Deirror/templette/props/aria"
	"github.com/Deirror/templette/props/data"

	"github.com/a-h/templ"
)

// Props wraps props.Attrs and embeds aria.Props and data.Props to provide
// a single container for general, ARIA, and data-* attributes for any HTML element.
type Props struct {
	props.Attrs

	Aria aria.Props
	Data data.Props
}

// AsTemplAttrs combines all underlying attributes and returns templ.Attributes
// for use in templ elements.
func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)
	return attrs.AsTemplAttrs()
}
