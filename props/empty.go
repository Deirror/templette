// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package props

import "github.com/a-h/templ"

// EmptyAttrsProvider is a no-op AttrsProvider.
// It can be used when you want to render an element without any attributes.
type EmptyAttrsProvider struct{}

// AsTemplAttrs returns an empty templ.Attributes map.
func (e EmptyAttrsProvider) AsTemplAttrs() templ.Attributes {
	return templ.Attributes{}
}
