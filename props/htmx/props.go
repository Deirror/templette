// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package htmx

import (
	"github.com/Deirror/templette/props"
)

// Props wraps props.Attrs to provide helpers for managing htmx (hx-*) attributes.
type Props struct {
	props.Attrs
}

// With sets a generic hx-* attribute with the given key and value.
func (h Props) With(key, val string) Props {
	h.Attrs = h.Attrs.With(HX+"-"+key, val)
	return h
}

// WithGet sets the "hx-get" attribute to the given URL.
func (h Props) WithGet(url string) Props {
	return h.withRequest(HXGet, url)
}

// WithPost sets the "hx-post" attribute to the given URL.
func (h Props) WithPost(url string) Props {
	return h.withRequest(HXPost, url)
}

// WithPut sets the "hx-put" attribute to the given URL.
func (h Props) WithPut(url string) Props {
	return h.withRequest(HXPut, url)
}

// WithPatch sets the "hx-patch" attribute to the given URL.
func (h Props) WithPatch(url string) Props {
	return h.withRequest(HXPatch, url)
}

// WithDelete sets the "hx-delete" attribute to the given URL.
func (h Props) WithDelete(url string) Props {
	return h.withRequest(HXDelete, url)
}

// WithTarget sets the "hx-target" attribute to the given selector.
func (h Props) WithTarget(sel string) Props {
	h.Attrs = h.Attrs.With(HXTarget, sel)
	return h
}

// WithTrigger sets the "hx-trigger" attribute to the given trigger string.
func (h Props) WithTrigger(trigger string) Props {
	h.Attrs = h.Attrs.With(HXTrigger, trigger)
	return h
}

// WithSwap sets the "hx-swap" attribute to the given swap strategy.
func (h Props) WithSwap(swap string) Props {
	h.Attrs = h.Attrs.With(HXSwap, swap)
	return h
}

// WithConfirm sets the "hx-confirm" attribute to the given message.
func (h Props) WithConfirm(msg string) Props {
	h.Attrs = h.Attrs.With(HXConfirm, msg)
	return h
}

// WithInclude sets the "hx-include" attribute to the given selector.
func (h Props) WithInclude(sel string) Props {
	h.Attrs = h.Attrs.With(HXInclude, sel)
	return h
}

// WithIndicator sets the "hx-indicator" attribute to the given selector.
func (h Props) WithIndicator(sel string) Props {
	h.Attrs = h.Attrs.With(HXIndicator, sel)
	return h
}

// clearRequest removes all hx-* request attributes (get, post, put, patch, delete).
func (h Props) clearRequest() {
	delete(h.Attributes, HXGet)
	delete(h.Attributes, HXPost)
	delete(h.Attributes, HXPut)
	delete(h.Attributes, HXPatch)
	delete(h.Attributes, HXDelete)
}

// withRequest clears existing hx-* request attributes and sets the given method to the URL.
// Returns a new Props with the attribute set.
func (h Props) withRequest(method, url string) Props {
	h.clearRequest()

	h.Attrs = h.Attrs.With(method, url)
	return h
}
