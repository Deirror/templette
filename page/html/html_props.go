package html

import "github.com/Deirror/templette/props"

type Props = props.Attrs

func NewHTML() Props {
	return props.NewAttrs()
}
