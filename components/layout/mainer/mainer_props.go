package mainer // Renamed due to 'main' being reserved

import (
	"github.com/Deirror/templette/props"
)

type Props = props.Attrs

func NewMain() Props {
	return props.NewAttrs()
}
