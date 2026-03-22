// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package props

// FirstOrDefault returns the first value from props if any are provided.
// Otherwise, it returns the default value def.
func FirstOrDefault[T any](def T, props ...T) T {
    if len(props) > 0 {
        return props[0]
    }
    return def
}
