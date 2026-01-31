package head

// Props represents <head> configuration
type Props struct {
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
