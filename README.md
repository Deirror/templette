# templette

`templette` is a Go front-end toolkit built on top of [templ](https://templ.guide). It focuses on reusable server-rendered HTML building blocks: typed HTML attributes, dynamic HTML element helpers, page wrappers, reusable component packages, and higher-level wrappers for composing application UI.

## What this repository provides

- **Reusable HTML components** in `components/`
- **Dynamic tag rendering** in `element/`
- **Page-level layout helpers** in `page/`
- **Typed attribute helpers** in `props/`
- **Reusable higher-level wrappers** in `wrappers/`

The goal is to let you build `.templ` templates from small, composable primitives instead of writing raw HTML everywhere.

## Package overview

### `components`
A collection of reusable HTML component packages. Each subpackage focuses on a specific HTML concern or component family, such as:

- `button`
- `div`
- `dropdown`
- `fieldset`
- `form`
- `input`
- `label`
- `layout`
- `media`
- `meter`
- `progress`
- `semantic`
- `table`
- `text`
- `textarea`

These packages are meant to be imported and used as ready-made building blocks inside your templ components.

### `element`
A low-level HTML element package for dynamic tag creation.

It exposes:

- `Tag` — a typed list of supported HTML tags
- `Element(t Tag, ...props.AttrsProvider)` — renders a normal HTML element
- `VoidElement(t Tag, ...props.AttrsProvider)` — renders a self-closing / void element

This package is useful when you want one helper that can render many tags at runtime while still staying inside the templ component model.

### `props`
A typed attribute layer around `templ.Attributes`.

It provides:

- `Attrs` — a wrapper around `templ.Attributes`
- `With(key, value)` — set any attribute
- `WithID(id)` — set the `id`
- `WithClass(class)` — append or set CSS classes
- `WithStyle(style)` — append or set inline styles
- `Merge(other)` — merge attribute maps
- `AsTemplAttrs()` — convert to `templ.Attributes`
- `AttrsProvider` — an interface for anything that can produce templ attributes

This package is the foundation for passing HTML attributes through the rest of the toolkit.

### `page`
Page-level wrappers for full HTML documents.

The package centers around:

- `Props` — a struct that groups:
  - `HTML`
  - `Head`
  - `Body`

Use this package when you need a full page shell rather than a single component.

### `wrappers`
Reusable high-level wrappers built on top of the lower-level component packages.

These wrappers are intended for consistent UI composition and shared styling. The package-level props are mutable, which makes global customization easy, but they are **not thread-safe** to change at runtime.

## How the pieces fit together

A typical flow looks like this:

1. Build attributes with `props.Attrs` or wrappers around the struct.
2. Render a tag with `element.Element` or `element.VoidElement`
3. Compose those tags into reusable components in `components/`
4. Wrap whole pages with `page/`
5. Apply shared styles and patterns with `wrappers/`

## Example

```go
attrs := props.Attrs{}.
    WithID("hero").
    WithClass("container").
    With("data-role", "hero")

card := element.Element(element.Div, attrs)
logo := element.VoidElement(element.Img, props.Attrs{}.
    With("src", "/logo.svg").
    With("alt", "Logo"))
```

```go
import (
    "github.com/Deirror/templette/components/div"
    "github.com/Deirror/templette/components/media/img"
)

attrs := div.Props{
    Attrs: props.Attrs{}.
        WithID("hero").
        WithClass("container"),
    Data: data.Props{}.With("role", "hero")
}

card := div.Div(attrs) // uses element.Element templ func
```

Inside `.templ` files, these helpers can be combined to build structured HTML without dropping down to raw markup for every element.

## Requirements

- Go `1.25.4`
- `templ` `v0.3.960`

## Notes

- The generated files in this repository are produced by `templ` and should not be edited by hand.
- `templette` is designed for server-rendered HTML and templ-based composition.
- The project is MIT licensed.

## Contributing

Issues and pull requests are welcome.
