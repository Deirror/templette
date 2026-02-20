package aria

import (
	"github.com/Deirror/templette/props"
)

type Props struct {
	props.Attrs
}

func (p Props) With(key, val string) Props {
	p.Attrs = p.Attrs.With("aria-"+key, val)
	return p
}
