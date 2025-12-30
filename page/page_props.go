package page

import (
	"github.com/Deirror/templette/props"
)

// HTMLProps represents attributes for the <html> tag
type HTMLProps struct {
	Attrs props.Attrs // lang, dir and etc.
}

// HeadProps represents <head> configuration
type HeadProps struct {
	Title   string    // <title>
	CSS     []string  // <link rel="stylesheet" href=...>
	Scripts []string  // <script src=...></script>
	Icons   []string  // <link rel="icon" href=...>
	Fonts   []string  // <link rel="stylesheet" href=...> for fonts
	Meta    []MetaTag // additional <meta> tags
	UseHTMX bool      // include htmx
}

// Optional helper for meta tags
type MetaTag struct {
	Name    string
	Content string
}

// BodyProps represents <body> configuration
type BodyProps struct {
	Attrs props.Attrs
	// Children []Component // all components rendered inside <body>
}

// PageProps represents the entire page
type PageProps struct {
	HTML HTMLProps
	Head HeadProps
	Body BodyProps
}
