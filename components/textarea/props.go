// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package textarea

import (
	"strconv"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"

	"github.com/Deirror/templette/props/aria"
	"github.com/Deirror/templette/props/data"
	"github.com/Deirror/templette/props/htmx"
)

type Props struct {
	props.Attrs

	Hx   htmx.Props
	Aria aria.Props
	Data data.Props

	Name        string
	Placeholder string
	Rows        int
	Cols        int

	Disabled bool
	Required bool
	ReadOnly bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)
	
	if p.Name != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Name, p.Name))
	}
	if p.Placeholder != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Placeholder, p.Placeholder))
	}
	if p.Rows > 0 {
		attrs = attrs.Merge(props.Attrs{}.With(props.Rows, strconv.Itoa(p.Rows)))
	}
	if p.Cols > 0 {
		attrs = attrs.Merge(props.Attrs{}.With(props.Cols, strconv.Itoa(p.Cols)))
	}
	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{}.With(props.Disabled, ""))
	}
	if p.Required {
		attrs = attrs.Merge(props.Attrs{}.With(props.Required, ""))
	}
	if p.ReadOnly {
		attrs = attrs.Merge(props.Attrs{}.With(props.ReadOnly, ""))
	}

	return attrs.AsTemplAttrs()
}
