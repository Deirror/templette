package props

type AriaProps struct {
	Attrs
}

func NewAria() AriaProps {
	return AriaProps{Attrs: NewAttrs()}
}

func (p AriaProps) With(key, val string) AriaProps {
	p.Attributes["aria-"+key] = val
	return p
}
