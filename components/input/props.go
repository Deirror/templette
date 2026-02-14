package input

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"

	"github.com/Deirror/templette/props/aria"
	"github.com/Deirror/templette/props/data"
	"github.com/Deirror/templette/props/htmx"
)

type Props struct {
	props.Attrs

	Hx   htmx.Props
	Aria aria.Props
	Data data.Props

	Type        string
	Name        string
	Value       string
	Placeholder string

	Disabled bool
	Required bool
	ReadOnly bool
	Checked  bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}

	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	if p.Type != "" {
		attrs = attrs.Merge(props.Attrs{}.With("type", p.Type))
	}
	if p.Name != "" {
		attrs = attrs.Merge(props.Attrs{}.With("name", p.Name))
	}
	if p.Value != "" {
		attrs = attrs.Merge(props.Attrs{}.With("value", p.Value))
	}
	if p.Placeholder != "" {
		attrs = attrs.Merge(props.Attrs{}.With("placeholder", p.Placeholder))
	}

	// boolean attributes
	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{}.With("disabled", ""))
	}
	if p.Required {
		attrs = attrs.Merge(props.Attrs{}.With("required", ""))
	}
	if p.ReadOnly {
		attrs = attrs.Merge(props.Attrs{}.With("readonly", ""))
	}
	if p.Checked {
		attrs = attrs.Merge(props.Attrs{}.With("checked", ""))
	}

	return attrs.AsTemplAttrs()
}

