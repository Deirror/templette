package form

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Hx   props.HxProps
	Aria props.AriaProps
	Data props.DataProps

	Action  string
	Method  string
	EncType string
	Target  string
	NoValid bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}

	attrs = attrs.Merge(p.Attrs)

	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	if p.Action != "" {
		attrs = attrs.Merge(props.Attrs{}.With("action", p.Action))
	}

	if p.Method != "" {
		attrs = attrs.Merge(props.Attrs{}.With("method", p.Method))
	}

	if p.EncType != "" {
		attrs = attrs.Merge(props.Attrs{}.With("enctype", p.EncType))
	}

	if p.Target != "" {
		attrs = attrs.Merge(props.Attrs{}.With("target", p.Target))
	}

	if p.NoValid {
		attrs = attrs.Merge(props.Attrs{}.With("novalidate", ""))
	}

	return attrs.AsTemplAttrs()
}

