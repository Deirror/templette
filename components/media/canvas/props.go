// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package canvas

import (
	"strconv"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Width  int
	Height int
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)

	if p.Width > 0 {
		attrs = attrs.Merge(props.Attrs{}.With(props.Width, strconv.Itoa(p.Width)))
	}
	if p.Height > 0 {
		attrs = attrs.Merge(props.Attrs{}.With(props.Height, strconv.Itoa(p.Height)))
	}

	return attrs.AsTemplAttrs()
}
