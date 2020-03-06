// Package ftmpl is generated with ftmpl {{{v0.3.1}}}, do not edit!!!! */
package ftmpl

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"os"

	"github.com/qorpress/benchmarks/go-template/model"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// TMPLERRindex2 evaluates a template index2.tmpl
func TMPLERRindex2(u *model.User, nav []*model.Navigation, title string) (string, error) {
	_template := "index2.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`
`)
	_w(`<!DOCTYPE html>
<html>
<body>

<header>
`)
	_w(`<title>`)
	_w(fmt.Sprintf(`%s`, _escape(title)))
	_w(`'s Home Page</title>
<div class="header">Page Header</div>`)
	_w(`

`)
	_w(`</header>

<nav>
`)
	_w(`<ul class="navigation">
    `)
	for _, item := range nav {
		_w(`
        	<li><a href="`)
		_w(fmt.Sprintf(`%s`, _escape(item.Link)))
		_w(`">`)
		_w(fmt.Sprintf(`%s`, _escape(item.Item)))
		_w(`</a></li>
    `)
	}
	_w(`
</ul>`)
	_w(`

`)
	_w(`</nav>

<section>
`)
	_w(`<div class="content">
	<div class="welcome">
		<h4>Hello `)
	_w(fmt.Sprintf(`%s`, _escape(u.FirstName)))
	_w(`</h4>
		
		<div class="raw">`)
	_w(fmt.Sprintf(`%s`, u.RawContent))
	_w(`</div>
		<div class="enc">`)
	_w(fmt.Sprintf(`%s`, _escape(u.EscapedContent)))
	_w(`</div>
	</div>


`)
	for i := 1; i <= 5; i++ {
		if i == 1 {
			_w(`			<p>`)
			_w(fmt.Sprintf(`%s`, _escape(u.FirstName)))
			_w(` has `)
			_w(fmt.Sprintf(`%d`, i))
			_w(` message</p>
`)
		} else {
			_w(`			<p>`)
			_w(fmt.Sprintf(`%s`, _escape(u.FirstName)))
			_w(` has `)
			_w(fmt.Sprintf(`%d`, i))
			_w(` messages</p>
`)
		}
	}
	_w(`</div>
`)
	_w(`</section>

<footer>
`)
	_w(`<div class="footer">copyright 2016</div>`)
	_w(`

`)
	_w(`</footer>

</body>
</html>`)

	return _ftmpl.String(), nil
}

// TMPLindex2 evaluates a template index2.tmpl
func TMPLindex2(u *model.User, nav []*model.Navigation, title string) string {
	html, err := TMPLERRindex2(u, nav, title)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template index2.tmpl:" + err.Error())
	}
	return html
}
