package models

import (
	"log"
	"net/http"

	db "github.com/JIeeiroSst/config"
	session "github.com/JIeeiroSst/utils"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id       int
	Name     string
	Password string
}

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login", nil)
}

func ShowSignUp(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "signup", nil)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	redirectTarget := "/admin"
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		result, err := db.Query("select password from admin where name=?", name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		admin := Admin{}
		for result.Next() {
			var passwords string
			if err = result.Scan(&passwords); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			admin.Password = passwords
			if err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				redirectTarget = "/admin/bug"
			}
		}
		session.SetSession(name, w)
	}
	defer db.Close()
	http.Redirect(w, r, redirectTarget, 301)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		log.Println(name + "|" + password)
		insert, err := db.Prepare("insert into Admin(name,password) values(?,?)")
		if err != nil {
			panic(err.Error())
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
		log.Println(hashedPassword)
		log.Println(string(hashedPassword))
		insert.Exec(name, string(hashedPassword))
	}
	defer db.Close()
	http.Redirect(w, r, "/admin/login", 301)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		session.ClearSession(w)
		http.Redirect(w, r, "/admin/login", 302)
	}
}
