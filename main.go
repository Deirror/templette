package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/Deirror/templette/components/button"
	"github.com/Deirror/templette/page"
	"github.com/Deirror/templette/props"

	"github.com/a-h/templ"
)

type Text string

func (t Text) Render(ctx context.Context, w io.Writer) error {
	_, err := io.WriteString(w, string(t))
	return err
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Build the btn props
		btn := button.Props{
			Attrs: props.NewAttrs().WithClass("btn").WithID("save-btn"),
			Type:  "button",
		}

		// Build the page props
		p := page.PageProps{
			HTML: page.HTMLProps{
				Attrs: props.NewAttrs().With("lang", "en"),
			},
			Head: page.HeadProps{
				Title:   "Home",
				CSS:     []string{"/static/btn.css"},
				Scripts: []string{},
				Fonts:   []string{},
				Icons:   []string{},
				Meta:    []page.MetaTag{},
				UseHTMX: true,
			},
			Body: page.BodyProps{
				Attrs: props.NewAttrs(),
			},
		}

		btnCtx := templ.WithChildren(context.Background(),
			Text("Save"), // child of the button
		)

		pageCtx := templ.WithChildren(
			context.Background(),
			templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				return button.Button(btn).Render(btnCtx, w)
			}),
		)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page.Page(p).Render(pageCtx, w) // call your generated templ func for the whole page
	})

	// Serve static files from ./static
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
