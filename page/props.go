package page

import (
	"github.com/Deirror/templette/page/body"
	"github.com/Deirror/templette/page/head"
	"github.com/Deirror/templette/page/html"
)

// PageProps represents the entire page
type Props struct {
	HTML html.Props
	Head head.Props
	Body body.Props
}
