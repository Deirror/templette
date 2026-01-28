# templette

Description
-

Conveniet and aimed at ease of use front-end UI toolkit, using Go's HTML templating lib(Templ)

Consists of three parts:
- `props` - defines ids, classes and other attrs for the given template
- `page` - website layout
- `components` - directly compiled templ Go files ready to use. Contains most used html tags like <div>, <text>, <p> and many, many other

The idea is that the developer can directly write `.templ` using the ready templ funcs and define props like htmx, data, aria and others

No JS at all, if you want you can write some scripts - there is a ready html tag which just needs a path to set from the page props

## Contributing

Contributions are welcome 🤝! Please open an issue or pull request
