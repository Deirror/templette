// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package overlay

import (
	"github.com/Deirror/templette/props"
)

var (
	Props props.Attrs = props.Attrs{}.WithStyle("padding: 1rem; background: rgba(0, 0, 0, 0.2); top: 0; left: 0; width: 100%; height: 100%; pointer-events: none") // default props attribute

	SoftProps  props.Attrs = props.Attrs{}.WithStyle("padding: 1rem; background: rgba(0, 0, 0, 0.1); top: 0; left: 0; width: 100%; height: 100%; pointer-events: none")
	DarkProps  props.Attrs = props.Attrs{}.WithStyle("padding: 1rem; background: rgba(0, 0, 0, 0.3); top: 0; left: 0; width: 100%; height: 100%; pointer-events: none")
	XDarkProps props.Attrs = props.Attrs{}.WithStyle("padding: 1rem; background: rgba(0, 0, 0, 0.4); top: 0; left: 0; width: 100%; height: 100%; pointer-events: none")
)
