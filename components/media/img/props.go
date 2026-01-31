package img

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Src string
	Alt string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.NewAttrs().Merge(p.Attrs)
	if p.Src != "" {
		attrs = attrs.Merge(props.NewAttrs().With("src", p.Src))
	}
	if p.Alt != "" {
		attrs = attrs.Merge(props.NewAttrs().With("alt", p.Alt))
	}
	return attrs.AsTemplAttrs()
}

