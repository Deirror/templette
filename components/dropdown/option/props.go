package option

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Value    string
	Selected bool
	Disabled bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}
	attrs = attrs.Merge(p.Attrs)

	if p.Value != "" {
		attrs = attrs.Merge(props.Attrs{}.With("value", p.Value))
	}
	if p.Selected {
		attrs = attrs.Merge(props.Attrs{}.With("selected", ""))
	}
	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{}.With("disabled", ""))
	}

	return attrs.AsTemplAttrs()
}
