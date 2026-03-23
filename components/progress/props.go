// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package progress

import (
	"fmt"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Value float64 // current value
	Max   float64 // maximum value
	Label string  // optional for accessibility
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(props.Attrs{}.
		With(props.Value, fmt.Sprintf("%g", p.Value)).
		With(props.Max, fmt.Sprintf("%g", p.Max)))

	if p.Label != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.AriaLabel, p.Label))
	}

	return attrs.AsTemplAttrs()
}
