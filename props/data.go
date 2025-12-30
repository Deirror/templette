package props

type DataProps map[string]string

func (d DataProps) WithAttr(key, val string) DataProps {
	d[key] = val
	return d
}

func (d DataProps) Attrs() Attrs {
	if len(d) == 0 {
		return nil
	}

	attrs := Attrs{}
	for k, v := range d {
		if len(k) > 0 && k[:5] != "data-" {
			attrs["data-"+k] = v
		} else {
			attrs[k] = v
		}
	}

	return attrs
}
