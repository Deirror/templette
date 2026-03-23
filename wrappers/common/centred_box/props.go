// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package centred_box

import (
	"github.com/Deirror/templette/props"
)

var (
	Props props.Attrs = props.Attrs{}.WithStyle("max-width: 80%; margin: 2rem auto; text-align: center") // default props attributes

	FullWidthProps props.Attrs = props.Attrs{}.WithStyle("max-width: 100%; margin: 2rem auto; text-align: center") // full-width centred box
)
