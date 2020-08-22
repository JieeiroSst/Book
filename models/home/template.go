package home

import "html/template"

var tmpl = template.Must(template.ParseGlob("views/admin/*"))
