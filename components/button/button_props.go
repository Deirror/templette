package button

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Hx   *props.HxProps
	Aria *props.AriaProps
	Data *props.DataProps

	Disabled bool
	Href     string
	Type     string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.NewAttrs()

	attrs = attrs.Merge(p.Attrs)

	if p.Aria != nil {
		attrs = attrs.Merge(p.Aria.Attrs)
	}
	if p.Hx != nil {
		attrs = attrs.Merge(p.Hx.Attrs)
	}
	if p.Data != nil {
		attrs = attrs.Merge(p.Data.Attrs)
	}

	if p.Disabled {
		attrs = attrs.Merge(props.NewAttrs().
			With("disabled", ""))
	}
	if p.Type != "" {
		attrs = attrs.Merge(props.NewAttrs().
			With("type", p.Type))
	}
	if p.Href != "" {
		attrs = attrs.Merge(props.NewAttrs().
			With("href", p.Href))
	}

	return attrs.AsTemplAttrs()
}
