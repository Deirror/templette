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
	attrs := props.Attrs{}.Merge(p.Attrs)
	if p.Src != "" {
		attrs = attrs.Merge(props.Attrs{}.With("src", p.Src))
	}
	if p.Alt != "" {
		attrs = attrs.Merge(props.Attrs{}.With("alt", p.Alt))
	}
	return attrs.AsTemplAttrs()
}

