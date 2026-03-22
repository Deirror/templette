package container

import (
	"github.com/Deirror/templette/components/card"
	"github.com/Deirror/templette/components/semantic/nav"
	"github.com/Deirror/templette/props"
)

var cardProps card.Props = card.Props{
	Attrs: props.Attrs{}.WithClass("container"),
}

var navProps nav.Props = nav.Props{
	Attrs: props.Attrs{}.WithClass("container"),
}
