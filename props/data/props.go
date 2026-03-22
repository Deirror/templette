// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package data

import (
	"github.com/Deirror/templette/props"
)

// Props wraps props.Attrs to provide helpers for managing HTML data-* attributes.
type Props struct {
	props.Attrs
}

// With adds a data-* attribute with the given key and value.
// The key is automatically prefixed with "data-".
func (p Props) With(key, val string) Props {
	p.Attrs = p.Attrs.With(Data+"-"+key, val)
	return p
}
