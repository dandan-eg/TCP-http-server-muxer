package template

import (
	"fmt"
	"strings"
)

type TemplateFunc func(param map[string]string) string

func Home(_ map[string]string) string {
	sb := strings.Builder{}

	sb.WriteString("<h1>Home</home></h1>")

	sb.WriteString("<p><a href=\"register\">Register</a></p>")

	return sb.String()
}

func Account(param map[string]string) string {

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("<p>Logged as %s</p>", param["name"]))
	sb.WriteString("<form method=\"post\" action=\"account\" >")
	sb.WriteString("<input type=\"submit\" value=\"delete\"/>")
	sb.WriteString("</form>")

	return sb.String()
}

func Delete(_ map[string]string) string {
	sb := strings.Builder{}

	sb.WriteString("<p>account deleted</p>")
	sb.WriteString("<p><a href=\"/\">back to index</a></p>")

	return sb.String()
}

func Register(_ map[string]string) string {
	sb := strings.Builder{}

	sb.WriteString(`<form method="get" action="/account"  >`)
	sb.WriteString(`<label for="name">Name :</label>`)
	sb.WriteString(`<input id="name" name="name" type="text"/>`)
	sb.WriteString(`<input type="submit" value="register"/>`)
	sb.WriteString(`</form>`)

	return sb.String()
}

func NotFound(_ map[string]string) string {

	return "<h1>404 not found :-(<h1/>"
}
