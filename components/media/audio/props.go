// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package audio

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Src      string
	Controls bool
	Autoplay bool
	Loop     bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)

	if p.Src != "" {
		attrs = attrs.Merge(props.Attrs{}.With(props.Src, p.Src))
	}
	if p.Controls {
		attrs = attrs.Merge(props.Attrs{}.With(props.Controls, ""))
	}
	if p.Autoplay {
		attrs = attrs.Merge(props.Attrs{}.With(props.Autoplay, ""))
	}
	if p.Loop {
		attrs = attrs.Merge(props.Attrs{}.With(props.Loop, ""))
	}

	return attrs.AsTemplAttrs()
}
