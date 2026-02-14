package props

type DataProps struct {
	Attrs
}

func (p DataProps) With(key, val string) DataProps {
	p.Attrs.With("data-"+key, val)
	return p
}
