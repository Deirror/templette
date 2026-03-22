// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package page

import (
	"github.com/Deirror/templette/page/body"
	"github.com/Deirror/templette/page/head"
	"github.com/Deirror/templette/page/html"
)

// Props represents the entire page
type Props struct {
	HTML html.Props
	Head head.Props
	Body body.Props
}
