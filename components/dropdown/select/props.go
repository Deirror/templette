// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package selectx

import (
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

	Name     string
	Multiple bool
	Required bool
	Disabled bool
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
	if p.Multiple {
		attrs = attrs.Merge(props.Attrs{}.With(props.Multiple, ""))
	}
	if p.Required {
		attrs = attrs.Merge(props.Attrs{}.With(props.Required, ""))
	}
	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{}.With(props.Disabled, ""))
	}

	return attrs.AsTemplAttrs()
}
