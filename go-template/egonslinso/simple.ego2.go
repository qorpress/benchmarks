// Generated by egon.
// 🚫Edit at your own risk.

package egonslinso

import (
	"html"
	"io"

	"github.com/qorpress/benchmarks/go-template/model"
)

func SimpleTemplate(w io.Writer, u *model.User) error {
	io.WriteString(w, "\n")
	io.WriteString(w, "\n<html>\n    <body>\n        <h1>")
	io.WriteString(w, html.EscapeString(u.FirstName))
	io.WriteString(w, "</h1>\n\n        <p>Here's a list of your favorite colors:</p>\n        <ul>\n        ")
	for _, colorName := range u.FavoriteColors {
		io.WriteString(w, "\n            <li>")
		io.WriteString(w, html.EscapeString(colorName))
		io.WriteString(w, "</li>")
	}
	io.WriteString(w, "\n        </ul>\n    </body>\n</html>")
	return nil
}
