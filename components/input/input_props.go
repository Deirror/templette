package input

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Hx   *props.HxProps
	Aria *props.AriaProps
	Data *props.DataProps

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

	if p.Type != "" {
		attrs = attrs.Merge(props.NewAttrs().With("type", p.Type))
	}
	if p.Name != "" {
		attrs = attrs.Merge(props.NewAttrs().With("name", p.Name))
	}
	if p.Value != "" {
		attrs = attrs.Merge(props.NewAttrs().With("value", p.Value))
	}
	if p.Placeholder != "" {
		attrs = attrs.Merge(props.NewAttrs().With("placeholder", p.Placeholder))
	}

	// boolean attributes
	if p.Disabled {
		attrs = attrs.Merge(props.NewAttrs().With("disabled", ""))
	}
	if p.Required {
		attrs = attrs.Merge(props.NewAttrs().With("required", ""))
	}
	if p.ReadOnly {
		attrs = attrs.Merge(props.NewAttrs().With("readonly", ""))
	}
	if p.Checked {
		attrs = attrs.Merge(props.NewAttrs().With("checked", ""))
	}

	return attrs.AsTemplAttrs()
}

