// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package option

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Value    string
	Selected bool
	Disabled bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)

	if p.Value != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Value, p.Value))
	}
	if p.Selected {
		attrs = attrs.Merge(props.Attrs{}.With(props.Selected, ""))
	}
	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{}.With(props.Disabled, ""))
	}

	return attrs.AsTemplAttrs()
}
