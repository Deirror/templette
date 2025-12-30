package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Deirror/templette/components/button"
	"github.com/Deirror/templette/page"
	"github.com/Deirror/templette/props"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Build the page props
		p := page.PageProps{
			HTML: page.HTMLProps{
				Attrs: props.NewAttrs().With("lang", "en"),
			},
			Head: page.HeadProps{
				Title:   "Home",
				CSS:     []string{"/app.css"},
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

		btn := button.Props{
			Attrs: props.NewAttrs().WithClass("btn-primary").WithID("save-btn"),
			Type:  "button",
		}

		page.Page(p).Render(context.Background(), w) // call your generated templ func for the whole page
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
