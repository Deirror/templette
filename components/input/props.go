// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package input

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

	Type        string
	Name        string
	Value       string
	Placeholder string
	Disabled    bool
	Required    bool
	ReadOnly    bool
	Checked     bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	if p.Type != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Type, p.Type))
	}
	if p.Name != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Name, p.Name))
	}
	if p.Value != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Value, p.Value))
	}
	if p.Placeholder != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Placeholder, p.Placeholder))
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
	if p.Checked {
		attrs = attrs.Merge(props.Attrs{}.With(props.Checked, ""))
	}

	return attrs.AsTemplAttrs()
}
