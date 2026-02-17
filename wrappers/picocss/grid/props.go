package grid

import (
	"github.com/Deirror/templette/components/card"
	"github.com/Deirror/templette/components/semantic/nav"
	"github.com/Deirror/templette/props"
)

var pCard card.Props = card.Props{
	Attrs: props.Attrs{}.WithClass("grid"),
}

var pNav nav.Props = nav.Props{
	Attrs: props.Attrs{}.WithClass("grid"),
}
