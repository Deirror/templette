// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package source

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Src  string
	Type string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)

	if p.Src != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Src, p.Src))
	}
	if p.Type != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Type, p.Type))
	}

	return attrs.AsTemplAttrs()
}
