package label

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Aria *props.AriaProps
	Data *props.DataProps

	For string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.NewAttrs()
	attrs = attrs.Merge(p.Attrs)

	if p.Aria != nil {
		attrs = attrs.Merge(p.Aria.Attrs)
	}
	if p.Data != nil {
		attrs = attrs.Merge(p.Data.Attrs)
	}

	if p.For != "" {
		attrs = attrs.Merge(
			props.NewAttrs().With("for", p.For),
		)
	}

	return attrs.AsTemplAttrs()
}

