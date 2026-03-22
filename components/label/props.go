// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package label

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"

	"github.com/Deirror/templette/props/aria"
	"github.com/Deirror/templette/props/data"
)

type Props struct {
	props.Attrs

	Aria aria.Props
	Data data.Props

	For string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)
	
	if p.For != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.For, p.For))
	}

	return attrs.AsTemplAttrs()
}

