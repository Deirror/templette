package meter

import (
	"fmt"

	"github.com/Deirror/templette/props"
	"github.com/a-h/templ"
)

type Props struct {
	props.Attrs

	Value float64 // current value
	Min   float64 // minimum
	Max   float64 // maximum
	Low   float64 // optional low threshold
	High  float64 // optional high threshold
	Opt   float64 // optional optimum value
	Label string  // aria-label
}

func (p Props) AsTemplAttrs() templ.Attributes {
	attrs := props.NewAttrs().Merge(p.Attrs)
	attrs = attrs.Merge(props.NewAttrs().
		With("value", fmt.Sprintf("%g", p.Value)).
		With("min", fmt.Sprintf("%g", p.Min)).
		With("max", fmt.Sprintf("%g", p.Max)))

	if p.Low != 0 {
		attrs = attrs.Merge(props.NewAttrs().With("low", fmt.Sprintf("%g", p.Low)))
	}
	if p.High != 0 {
		attrs = attrs.Merge(props.NewAttrs().With("high", fmt.Sprintf("%g", p.High)))
	}
	if p.Opt != 0 {
		attrs = attrs.Merge(props.NewAttrs().With("optimum", fmt.Sprintf("%g", p.Opt)))
	}
	if p.Label != "" {
		attrs = attrs.Merge(props.NewAttrs().With("aria-label", p.Label))
	}

	return attrs.AsTemplAttrs()
}

