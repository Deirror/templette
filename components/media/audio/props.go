package audio

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Src      string
	Controls bool
	Autoplay bool
	Loop     bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}.Merge(p.Attrs)
	if p.Src != "" {
		attrs = attrs.Merge(props.Attrs{}.With("src", p.Src))
	}
	if p.Controls {
		attrs = attrs.Merge(props.Attrs{}.With("controls", ""))
	}
	if p.Autoplay {
		attrs = attrs.Merge(props.Attrs{}.With("autoplay", ""))
	}
	if p.Loop {
		attrs = attrs.Merge(props.Attrs{}.With("loop", ""))
	}
	return attrs.AsTemplAttrs()
}

