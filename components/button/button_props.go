package button

import (
	"github.com/Deirror/templette/props"
)

type Props struct {
	props.Base

	Hx   *props.HxProps
	Aria *props.AriaProps
	Data *props.DataProps

	Disabled bool
	Href     string
	Type     string
}

func (p Props) AsAttrs() props.Attrs {
	attrs := p.Base.Attrs.
		Merge(p.Aria.Attrs()).
		Merge(p.Hx.Attrs()).
		Merge(p.Data.Attrs())

	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{"disabled": "disabled"})
	}
	if p.Type != "" {
		attrs = attrs.Merge(props.Attrs{"type": p.Type})
	}
	if p.Href != "" {
		attrs = attrs.Merge(props.Attrs{"href": p.Href})
	}

	return attrs
}
