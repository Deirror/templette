package group

import (
	"github.com/Deirror/templette/components/fieldset"
	"github.com/Deirror/templette/components/card"

	"github.com/Deirror/templette/props"
)

var fieldsetProps fieldset.Props = fieldset.Props{
	Attrs: props.Attrs{}.With("role","group"),
}

var cardProps card.Props = card.Props{
	Attrs: props.Attrs{}.With("role","group"),
}
