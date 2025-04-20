package server

import (
	"context"
	"net/http"

	"github.com/nosvagor/hgmx/components/button"
	"github.com/nosvagor/hgmx/components/index"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cmp := index.Index(index.Page{
		Title: "HGMX Builder",
		Content: button.Button(button.Props{
			Attrs: map[string]any{
				"onclick": "console.log('Button clicked')",
			},
			Text: "Click me",
		}),
	})

	cmp.Render(context.Background(), w)
}
