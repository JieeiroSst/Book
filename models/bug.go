package models

import (
	"net/http"
)

func Bug(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "bug", nil)
}
