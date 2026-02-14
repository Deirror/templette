package props

type AriaProps struct {
	Attrs
}

func (p AriaProps) With(key, val string) AriaProps {
	p.Attrs.With("aria-"+key, val)
	return p
}
