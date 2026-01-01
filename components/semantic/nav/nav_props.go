package nav

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs
}

func (p Props) AsTemplAttrs() templ.Attributes {
	return p.Attrs.AsTemplAttrs()
}
