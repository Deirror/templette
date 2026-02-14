package toast

import (
	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
	"time"
)

type Variant string

const (
	Success Variant = "success"
	Error   Variant = "error"
	Warning Variant = "warning"
	Info    Variant = "info"
)

type Props struct {
	props.Attrs

	Variant Variant      // success | error | info | warning
	Title   string       // optional heading
	Message string       // main content (fallback if children empty)
	Timeout time.Duration // auto-dismiss (0 = persistent)

	Hx   props.HxProps
	Aria props.AriaProps
	Data props.DataProps
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.Attrs{}.
		WithClass("toast").
		WithClass("toast--" + string(p.Variant))

	attrs = attrs.Merge(p.Attrs)

	if p.Timeout > 0 {
		attrs = attrs.With(
			"hx-trigger", "load delay:"+p.Timeout.String(),
		).With(
			"hx-swap", "delete",
		)
	}

	attrs = attrs.Merge(p.Aria.Attrs)
	attrs = attrs.Merge(p.Hx.Attrs)
	attrs = attrs.Merge(p.Data.Attrs)

	return attrs.AsTemplAttrs()
}

