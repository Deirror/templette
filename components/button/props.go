package button

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

	Disabled bool
	Href     string
	Type     string
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}

	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	if p.Disabled {
		attrs = attrs.Merge(props.Attrs{}.
			With("disabled", ""))
	}
	if p.Type != "" {
		attrs = attrs.Merge(props.Attrs{}.
			With("type", p.Type))
	}
	if p.Href != "" {
		attrs = attrs.Merge(props.Attrs{}.
			With("href", p.Href))
	}

	return attrs.AsTemplAttrs()
}
