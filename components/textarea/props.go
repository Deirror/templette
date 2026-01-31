package textarea

import (
	"strconv"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Hx   *props.HxProps
	Aria *props.AriaProps
	Data *props.DataProps

	Name        string
	Placeholder string
	Rows        int
	Cols        int

	Disabled bool
	Required bool
	ReadOnly bool
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
	if p.Placeholder != "" {
		attrs = attrs.Merge(props.NewAttrs().With("placeholder", p.Placeholder))
	}
	if p.Rows > 0 {
		attrs = attrs.Merge(props.NewAttrs().With("rows", strconv.Itoa(p.Rows)))
	}
	if p.Cols > 0 {
		attrs = attrs.Merge(props.NewAttrs().With("cols", strconv.Itoa(p.Cols)))
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

	return attrs.AsTemplAttrs()
}
