package source

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Src  string
	Type string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}.Merge(p.Attrs)

	if p.Src != "" {
		attrs = attrs.Merge(props.Attrs{}.With("src", p.Src))
	}
	if p.Type != "" {
		attrs = attrs.Merge(props.Attrs{}.With("type", p.Type))
	}

	return attrs.AsTemplAttrs()
}
