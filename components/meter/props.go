// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package meter

import (
	"fmt"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Value float64 // current value
	Min   float64 // minimum
	Max   float64 // maximum
	Low   float64 // optional low threshold
	High  float64 // optional high threshold
	Opt   float64 // optional optimum value
	Label string  // aria-label
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}.Merge(p.Attrs)
	attrs = attrs.Merge(props.Attrs{}.
		With(props.Value, fmt.Sprintf("%g", p.Value)).
		With(props.Min, fmt.Sprintf("%g", p.Min)).
		With(props.Max, fmt.Sprintf("%g", p.Max)))

	if p.Low != 0 {
		attrs = attrs.Merge(props.Attrs{}.With(props.Low, fmt.Sprintf("%g", p.Low)))
	}
	if p.High != 0 {
		attrs = attrs.Merge(props.Attrs{}.With(props.High, fmt.Sprintf("%g", p.High)))
	}
	if p.Opt != 0 {
		attrs = attrs.Merge(props.Attrs{}.With(props.Optimum, fmt.Sprintf("%g", p.Opt)))
	}
	if p.Label != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.AriaLabel, p.Label))
	}

	return attrs.AsTemplAttrs()
}

