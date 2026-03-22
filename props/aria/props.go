// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package aria

import (
	"github.com/Deirror/templette/props"
)

// Props wraps props.Attrs to provide helpers for managing ARIA attributes.
type Props struct {
	props.Attrs
}

// With adds an ARIA attribute with the given key and value.
// The key is automatically prefixed with "aria-".
func (p Props) With(key, val string) Props {
	p.Attrs = p.Attrs.With(Aria+"-"+key, val)
	return p
}
