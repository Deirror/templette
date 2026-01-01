package field

import (
	"github.com/Deirror/templette/props"
)

type Props struct {
	props.Attrs

	Aria *props.AriaProps
	Data *props.DataProps

	Label string
}

