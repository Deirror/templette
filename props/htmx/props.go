package htmx

import (
	"github.com/Deirror/templette/props"
)

type Props struct {
	props.Attrs
}

func (h Props) WithGet(url string) Props {
	return h.withRequest("hx-get", url)
}

func (h Props) WithPost(url string) Props {
	return h.withRequest("hx-post", url)
}

func (h Props) WithPut(url string) Props {
	return h.withRequest("hx-put", url)
}

func (h Props) WithPatch(url string) Props {
	return h.withRequest("hx-patch", url)
}

func (h Props) WithDelete(url string) Props {
	return h.withRequest("hx-delete", url)
}

func (h Props) WithTarget(sel string) Props {
	h.Attrs = h.Attrs.With("hx-target", sel)
	return h
}

func (h Props) WithTrigger(trigger string) Props {
	h.Attrs = h.Attrs.With("hx-trigger", trigger)
	return h
}

func (h Props) WithSwap(swap string) Props {
	h.Attrs = h.Attrs.With("hx-swap", swap)
	return h
}

func (h Props) WithConfirm(msg string) Props {
	h.Attrs = h.Attrs.With("hx-confirm", msg)
	return h
}

func (h Props) WithInclude(sel string) Props {
	h.Attrs = h.Attrs.With("hx-include", sel)
	return h
}

func (h Props) WithIndicator(sel string) Props {
	h.Attrs = h.Attrs.With("hx-indicator", sel)
	return h
}

func (h Props) clearRequest() {
	delete(h.Attributes, "hx-get")
	delete(h.Attributes, "hx-post")
	delete(h.Attributes, "hx-put")
	delete(h.Attributes, "hx-patch")
	delete(h.Attributes, "hx-delete")
}

func (h Props) withRequest(method, url string) Props {
	h.clearRequest()
	h.Attrs = h.Attrs.With(method, url)
	return h
}
