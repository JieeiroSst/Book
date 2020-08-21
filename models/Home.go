package models

import (
	"net/http"

	session "github.com/JIeeiroSst/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "home", userName)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
}
