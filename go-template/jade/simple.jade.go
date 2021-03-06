// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	pool "github.com/valyala/bytebufferpool"

	"github.com/qorpress/benchmarks/go-template/model"
)

const (
	simple__0 = `<html><body><h1>`
	simple__1 = `</h1><p>Here's a list of your favorite colors:</p><ul>`
	simple__2 = `</ul></body></html>`
	simple__3 = `<li>`
	simple__4 = `</li>`
)

func Simple(u *model.User, buffer *pool.ByteBuffer) {

	buffer.WriteString(simple__0)
	WriteEscString(u.FirstName, buffer)
	buffer.WriteString(simple__1)

	for _, colorName := range u.FavoriteColors {
		buffer.WriteString(simple__3)
		WriteEscString(colorName, buffer)
		buffer.WriteString(simple__4)
	}
	buffer.WriteString(simple__2)

}
