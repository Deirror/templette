package progress

import (
	"fmt"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Value float64 // current value
	Max   float64 // maximum value
	Label string  // optional for accessibility
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}.Merge(p.Attrs)
	attrs = attrs.Merge(props.Attrs{}.
		With("value", fmt.Sprintf("%g", p.Value)).
		With("max", fmt.Sprintf("%g", p.Max)))

	if p.Label != "" {
		attrs = attrs.Merge(props.Attrs{}.With("aria-label", p.Label))
	}

	return attrs.AsTemplAttrs()
}
