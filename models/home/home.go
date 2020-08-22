package home

import (
	"net/http"

	session "github.com/JIeeiroSst/utils"
)

func HomeHome(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "HomeHome", userName)
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
}
