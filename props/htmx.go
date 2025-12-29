package props

import (
	"errors"
)

type HxProps struct {
	Get    *string
	Post   *string
	Put    *string
	Patch  *string
	Delete *string

	Target    *string
	Trigger   *string
	Swap      *string
	Confirm   *string
	Include   *string
	Indicator *string
}

func (hx *HxProps) Attrs() Attrs {
	if hx == nil {
		return nil
	}

	attrs := Attrs{}

	if method, url, ok := hx.Request(); ok {
		attrs[method] = url
	}

	if hx.Target != nil {
		attrs["hx-target"] = *hx.Target
	}
	if hx.Trigger != nil {
		attrs["hx-trigger"] = *hx.Trigger
	}
	if hx.Swap != nil {
		attrs["hx-swap"] = *hx.Swap
	}
	if hx.Confirm != nil {
		attrs["hx-confirm"] = *hx.Confirm
	}
	if hx.Include != nil {
		attrs["hx-include"] = *hx.Include
	}
	if hx.Indicator != nil {
		attrs["hx-indicator"] = *hx.Indicator
	}

	if len(attrs) == 0 {
		return nil
	}

	return attrs
}

func (hx *HxProps) Request() (method string, url string, ok bool) {
	if hx == nil {
		return "", "", false
	}

	switch {
	case hx.Get != nil:
		return "hx-get", *hx.Get, true
	case hx.Post != nil:
		return "hx-post", *hx.Post, true
	case hx.Put != nil:
		return "hx-put", *hx.Put, true
	case hx.Patch != nil:
		return "hx-patch", *hx.Patch, true
	case hx.Delete != nil:
		return "hx-delete", *hx.Delete, true
	default:
		return "", "", false
	}
}

func (hx *HxProps) Validate() error {
	count := 0
	if hx.Get != nil {
		count++
	}
	if hx.Post != nil {
		count++
	}
	if hx.Put != nil {
		count++
	}
	if hx.Patch != nil {
		count++
	}
	if hx.Delete != nil {
		count++
	}

	if count > 1 {
		return errors.New("only one hx request method may be set")
	}
	return nil
}
