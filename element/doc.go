// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package element provides type-safe representations of HTML tags
// and helper functions for working with generic elements in templated code.
//
// It defines the Tag type and a set of constants for common HTML tags,
// including semantic elements, form controls, media elements, and text elements.
// The IsValid method allows runtime validation of tag names.
//
// This package is intended for internal use by higher-level components
// or wrappers that generate HTML via templates, providing a consistent
// way to refer to HTML tags across your codebase.
package element
