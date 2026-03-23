// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package nav

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs
}

func (p Props) AsTemplAttrs() templ.Attributes {
	return p.Attrs.AsTemplAttrs()
}
