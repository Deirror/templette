package props

type DataProps struct {
	Attrs
}

func NewData() DataProps {
	return DataProps{Attrs: NewAttrs()}
}

func (p DataProps) With(key, val string) DataProps {
	p.Attributes["data-"+key] = val
	return p
}
