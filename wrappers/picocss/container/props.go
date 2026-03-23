// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package container

import (
	"github.com/Deirror/templette/props"
)

var (
	Props props.Attrs = props.Attrs{}.WithClass("container") // default container attributes.

	FluidProps props.Attrs = props.Attrs{}.WithClass("container-fluid") // full-width container
)
