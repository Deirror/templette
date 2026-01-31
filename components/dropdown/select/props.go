package select_

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Hx   *props.HxProps
	Aria *props.AriaProps
	Data *props.DataProps

	Name     string
	Multiple bool
	Required bool
	Disabled bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.NewAttrs()
	attrs = attrs.Merge(p.Attrs)

	if p.Hx != nil {
		attrs = attrs.Merge(p.Hx.Attrs)
	}
	if p.Aria != nil {
		attrs = attrs.Merge(p.Aria.Attrs)
	}
	if p.Data != nil {
		attrs = attrs.Merge(p.Data.Attrs)
	}

	if p.Name != "" {
		attrs = attrs.Merge(props.NewAttrs().With("name", p.Name))
	}

	if p.Multiple {
		attrs = attrs.Merge(props.NewAttrs().With("multiple", ""))
	}
	if p.Required {
		attrs = attrs.Merge(props.NewAttrs().With("required", ""))
	}
	if p.Disabled {
		attrs = attrs.Merge(props.NewAttrs().With("disabled", ""))
	}

	return attrs.AsTemplAttrs()
}

