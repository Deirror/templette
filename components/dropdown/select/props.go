package select_

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Hx   props.HxProps
	Aria props.AriaProps
	Data props.DataProps

	Name     string
	Multiple bool
	Required bool
	Disabled bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}

	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	if p.Name != "" {
		attrs = attrs.Merge(props.Attrs{}.With("name", p.Name))
	}

	if p.Multiple {
		attrs = attrs.Merge(props.Attrs{}.With("multiple", ""))
	}
	if p.Required {
		attrs = attrs.Merge(props.Attrs{}.With("required", ""))
	}
	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{}.With("disabled", ""))
	}

	return attrs.AsTemplAttrs()
}

