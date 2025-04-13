package views

import (
	"embed"
	"encoding/json"
	"html/template"
)

// ViewsFS Include the template file in the binary
//
//go:embed templates/**/*
var ViewsFS embed.FS

type Layout int

const (
	NoLayout Layout = iota + 1
	NavbarLayout
	SidebarLayout
)

// ParseTemplate parses a set of template files in a specific nested order.
// The `index.gohtml` file is always used as the base template, and depending on the `mainLayout` flag,
// either the `layout_navbar.gohtml` or `layout_sidebar.gohtml` is used.
func ParseTemplate(layout Layout, templates ...string) (*template.Template, error) {
	if layout == NavbarLayout {
		templates = append([]string{"templates/console/layout_navbar.gohtml"}, templates...)
	}
	if layout == SidebarLayout {
		templates = append([]string{"templates/console/layout_sidebar.gohtml"}, templates...)
	}

	templates = append([]string{"templates/console/index.gohtml"}, templates...)

	tmpl, err := template.New("template").Funcs(template.FuncMap{
		"json": func(v interface{}) (string, error) {
			bytes, err := json.Marshal(v)
			return string(bytes), err
		},
	}).ParseFS(ViewsFS, templates...)
	if err != nil {
		return nil, err
	}

	return tmpl, err
}
