package html

import (
	"github.com/Deirror/templette/props"
	"github.com/Deirror/templette/props/aria"
	"github.com/Deirror/templette/props/data"

	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Aria aria.Props
	Data data.Props
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}

	attrs = attrs.Merge(p.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	return attrs.AsTemplAttrs()
}
