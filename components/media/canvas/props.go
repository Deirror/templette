package canvas

import (
	"strconv"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Width  int
	Height int
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}.Merge(p.Attrs)

	if p.Width > 0 {
		attrs = attrs.Merge(props.Attrs{}.With("width", strconv.Itoa(p.Width)))
	}
	if p.Height > 0 {
		attrs = attrs.Merge(props.Attrs{}.With("height", strconv.Itoa(p.Height)))
	}

	return attrs.AsTemplAttrs()
}
