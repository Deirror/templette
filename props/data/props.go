package data

import (
	"github.com/Deirror/templette/props"
)

type Props struct {
	props.Attrs
}

func (p Props) With(key, val string) Props {
	p.Attrs.With("data-"+key, val)
	return p
}
