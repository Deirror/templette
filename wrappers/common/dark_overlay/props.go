package dark_overlay

import (
	"github.com/Deirror/templette/components/card"
	"github.com/Deirror/templette/props"
)

var darkProps card.Props = card.Props{
	Attrs: props.Attrs{}.WithClass("dark-overlay"),
}

var softDarkProps card.Props = card.Props{
	Attrs: props.Attrs{}.WithClass("soft-dark-overlay"),
}
