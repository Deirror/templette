package label

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Aria props.AriaProps
	Data props.DataProps

	For string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}

	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)
	
	if p.For != "" {
		attrs = attrs.Merge(
			props.Attrs{}.With("for", p.For),
		)
	}

	return attrs.AsTemplAttrs()
}

