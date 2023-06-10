package main

import (
	"TCP-server-http/muxer"
	"TCP-server-http/template"
	"fmt"
)

func routes() muxer.Muxer {
	m := muxer.NewMuxer()

	m.Add("GET", "/", template.Home)
	m.Add("GET", "/account", template.Account)
	m.Add("POST", "/account", template.Delete)
	m.Add("GET", "/register", template.Register)

	fmt.Println(m)
	return m
}
