// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package text_align

import (
	"github.com/Deirror/templette/props"
)

var (
	LeftProps   props.Attrs = props.Attrs{}.WithStyle("text-align: left")
	CenterProps props.Attrs = props.Attrs{}.WithStyle("text-align: center")
	RightProps  props.Attrs = props.Attrs{}.WithStyle("text-align: right")
)
