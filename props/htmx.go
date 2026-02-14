package props

type HxProps struct {
	Attrs
}

func (h HxProps) WithGet(url string) HxProps {
	return h.withRequest("hx-get", url)
}

func (h HxProps) WithPost(url string) HxProps {
	return h.withRequest("hx-post", url)
}

func (h HxProps) WithPut(url string) HxProps {
	return h.withRequest("hx-put", url)
}

func (h HxProps) WithPatch(url string) HxProps {
	return h.withRequest("hx-patch", url)
}

func (h HxProps) WithDelete(url string) HxProps {
	return h.withRequest("hx-delete", url)
}

func (h HxProps) WithTarget(sel string) HxProps {
	h.Attrs.With("hx-target", sel)
	return h
}

func (h HxProps) WithTrigger(trigger string) HxProps {
	h.Attrs.With("hx-trigger", trigger)
	return h
}

func (h HxProps) WithSwap(swap string) HxProps {
	h.Attrs.With("hx-swap", swap)
	return h
}

func (h HxProps) WithConfirm(msg string) HxProps {
	h.Attrs.With("hx-confirm", msg)
	return h
}

func (h HxProps) WithInclude(sel string) HxProps {
	h.Attrs.With("hx-include", sel)
	return h
}

func (h HxProps) WithIndicator(sel string) HxProps {
	h.Attrs.With("hx-indicator", sel)
	return h
}

func (h HxProps) clearRequest() {
	delete(h.Attributes, "hx-get")
	delete(h.Attributes, "hx-post")
	delete(h.Attributes, "hx-put")
	delete(h.Attributes, "hx-patch")
	delete(h.Attributes, "hx-delete")
}

func (h HxProps) withRequest(method, url string) HxProps {
	h.clearRequest()
	h.Attrs.With(method, url)
	return h
}
