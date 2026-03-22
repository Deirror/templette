// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package form

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

	Action  string
	Method  string
	EncType string
	Target  string
	NoValid bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	if p.Method != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Method, p.Method))
	}
	if p.EncType != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.EncType, p.EncType))
	}
	if p.Target != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Target, p.Target))
	}
	if p.NoValid {
		attrs = attrs.Merge(props.Attrs{}.With(props.NoValidate, ""))
	}

	return attrs.AsTemplAttrs()
}
