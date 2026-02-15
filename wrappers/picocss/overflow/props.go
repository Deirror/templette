package overflow

import (
	"github.com/Deirror/templette/components/card"
	"github.com/Deirror/templette/props"
)

var p card.Props = card.Props{
	Attrs: props.Attrs{}.WithClass("overflow-auto"),
}
