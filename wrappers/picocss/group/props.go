package group

import (
	"github.com/Deirror/templette/components/fieldset"
	"github.com/Deirror/templette/props"
)

var p fieldset.Props = fieldset.Props{
	Attrs: props.Attrs{}.With("role","group"),
}
