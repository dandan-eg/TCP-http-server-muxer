package muxer

import (
	"TCP-server-http/request"
	"TCP-server-http/template"
	"fmt"
	"strings"
)

type Muxer []Route

func NewMuxer() Muxer {
	return make(Muxer, 0, 0)
}

func (m *Muxer) Serve(req request.Request) string {
	for _, route := range *m {
		match := route.method == req.Method && req.URI == route.uri

		if match {

			return route.template(req.Param)
		}
	}

	return template.NotFound(nil)
}

func (m *Muxer) Add(method, uri string, tmpl template.TemplateFunc) {
	r := Route{
		method:   method,
		uri:      uri,
		template: tmpl,
	}

	*m = append(*m, r)
}

func (m Muxer) String() string {
	sb := strings.Builder{}
	sb.WriteString("available routes : \n")
	for _, route := range m {
		sb.WriteString(fmt.Sprintf("\t-[%s] %s\n", route.method, route.uri))

	}

	return sb.String()
}
