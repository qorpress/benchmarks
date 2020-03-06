// This file is automatically generated by qtc from "navigation.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line quicktemplate/navigation.qtpl:1
package quicktemplate

//line quicktemplate/navigation.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"

	"github.com/qorpress/benchmarks/go-template/model"

//line quicktemplate/navigation.qtpl:1
)

//line quicktemplate/navigation.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line quicktemplate/navigation.qtpl:2
func StreamNavigation(qw422016 *qt422016.Writer, nav []*model.Navigation) {
	//line quicktemplate/navigation.qtpl:2
	qw422016.N().S(`
<ul class="navigation">
`)
	//line quicktemplate/navigation.qtpl:4
	for _, item := range nav {
		//line quicktemplate/navigation.qtpl:4
		qw422016.N().S(`
	<li><a href="`)
		//line quicktemplate/navigation.qtpl:5
		qw422016.E().S(item.Link)
		//line quicktemplate/navigation.qtpl:5
		qw422016.N().S(`">`)
		//line quicktemplate/navigation.qtpl:5
		qw422016.E().S(item.Item)
		//line quicktemplate/navigation.qtpl:5
		qw422016.N().S(`</a></li>
`)
		//line quicktemplate/navigation.qtpl:6
	}
	//line quicktemplate/navigation.qtpl:6
	qw422016.N().S(`
</ul>
`)
//line quicktemplate/navigation.qtpl:8
}

//line quicktemplate/navigation.qtpl:8
func WriteNavigation(qq422016 qtio422016.Writer, nav []*model.Navigation) {
	//line quicktemplate/navigation.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line quicktemplate/navigation.qtpl:8
	StreamNavigation(qw422016, nav)
	//line quicktemplate/navigation.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line quicktemplate/navigation.qtpl:8
}

//line quicktemplate/navigation.qtpl:8
func Navigation(nav []*model.Navigation) string {
	//line quicktemplate/navigation.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
	//line quicktemplate/navigation.qtpl:8
	WriteNavigation(qb422016, nav)
	//line quicktemplate/navigation.qtpl:8
	qs422016 := string(qb422016.B)
	//line quicktemplate/navigation.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
	//line quicktemplate/navigation.qtpl:8
	return qs422016
//line quicktemplate/navigation.qtpl:8
}
