package muxer

import "TCP-server-http/template"

type Route struct {
	method   string
	uri      string
	template template.TemplateFunc
}
