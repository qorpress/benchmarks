// Generated by egon.
// 🚫Edit at your own risk.

package egonslinso

import (
	"html"
	"io"
)

func HeaderTemplate(w io.Writer, title string) error {
	io.WriteString(w, "\n<title>")
	io.WriteString(w, html.EscapeString(title))
	io.WriteString(w, "'s Home Page</title>\n<div class=\"header\">Page Header</div>")
	return nil
}
