package props

type AriaProps map[string]string

func (a AriaProps) WithAttr(key, val string) AriaProps {
	a[key] = val
	return a
}

func (a AriaProps) Attrs() Attrs {
	if len(a) == 0 {
		return nil
	}

	attrs := Attrs{}
	for k, v := range a {
		// all aria- attributes must start with "aria-"
		if len(k) > 0 && k[:5] != "aria-" {
			attrs["aria-"+k] = v
		} else {
			attrs[k] = v
		}
	}

	return attrs
}
