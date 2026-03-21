package centred_box

import (
	"github.com/Deirror/templette/components/card"
	"github.com/Deirror/templette/props"
)

var widthProps card.Props = card.Props{
	Attrs: props.Attrs{}.WithStyle("max-width: 80%; margin: 2rem auto; text-align: center"),
}

var fullWidthProps card.Props = card.Props{
	Attrs: props.Attrs{}.WithStyle("max-width: 100%; margin: 2rem auto; text-align: center"),
}
