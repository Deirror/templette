package form

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Hx   *props.HxProps
	Aria *props.AriaProps
	Data *props.DataProps

	Action  string
	Method  string
	EncType string
	Target  string
	NoValid bool
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.NewAttrs()
	attrs = attrs.Merge(p.Attrs)

	if p.Hx != nil {
		attrs = attrs.Merge(p.Hx.Attrs)
	}
	if p.Aria != nil {
		attrs = attrs.Merge(p.Aria.Attrs)
	}
	if p.Data != nil {
		attrs = attrs.Merge(p.Data.Attrs)
	}

	if p.Action != "" {
		attrs = attrs.Merge(props.NewAttrs().With("action", p.Action))
	}

	if p.Method != "" {
		attrs = attrs.Merge(props.NewAttrs().With("method", p.Method))
	}

	if p.EncType != "" {
		attrs = attrs.Merge(props.NewAttrs().With("enctype", p.EncType))
	}

	if p.Target != "" {
		attrs = attrs.Merge(props.NewAttrs().With("target", p.Target))
	}

	if p.NoValid {
		attrs = attrs.Merge(props.NewAttrs().With("novalidate", ""))
	}

	return attrs.AsTemplAttrs()
}

