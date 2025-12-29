package button

import (
	"github.com/Deirror/templette/components/props"
	"github.com/Deirror/templette/props"
)

type Props struct {
	cprops.Base
	cprops.Variant
	cprops.Size

	Hx   *props.HxProps
	Aria *props.AriaProps

	Disabled bool
	Href     string
	Type     string

	cprops.Content
}
