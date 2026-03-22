// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package props

import "github.com/a-h/templ"

// AttrsProvider prepares all attributes into one merged templ.Attributes
// which is used in the html tags.
type AttrsProvider interface {
	AsTemplAttrs() templ.Attributes
}
