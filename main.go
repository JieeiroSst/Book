package main

import (
	"log"
	"net/http"

	model "github.com/JIeeiroSst/models/admin"
	models "github.com/JIeeiroSst/models/home"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	log.Println("Server started on: http://localhost:8080")
	route.HandleFunc("/admin", model.Home)

	route.HandleFunc("/admin/login", model.ShowLogin)
	route.HandleFunc("/admin/logins", model.Signin)
	route.HandleFunc("/admin/signup", model.ShowSignUp)
	route.HandleFunc("/admin/signups", model.Signup)
	route.HandleFunc("/admin/logout", model.LogoutHandler)
	route.HandleFunc("/admin/bug", model.Bug)

	route.HandleFunc("/admin/employe", model.Index)
	route.HandleFunc("/admin/employe/show", model.Show)
	route.HandleFunc("/admin/employe/new", model.New)
	route.HandleFunc("/admin/employe/edit", model.Edit)
	route.HandleFunc("/admin/employe/insert", model.Insert)
	route.HandleFunc("/admin/employe/update", model.Update)
	route.HandleFunc("/admin/employe/delete", model.Delete)

	route.HandleFunc("/admin/book", model.IndexBook)
	route.HandleFunc("/admin/book/show", model.ShowBook)
	route.HandleFunc("/admin/book/new", model.NewBook)
	route.HandleFunc("/admin/book/edit", model.EditBook)
	route.HandleFunc("/admin/book/insert", model.InsertBook)
	route.HandleFunc("/admin/book/delete", model.DeleteBook)

	route.HandleFunc("/admin/user", model.IndexUser)

	route.HandleFunc("/admin/category", model.IndexCategory)
	route.HandleFunc("/admin/category/show", model.ShowCattegory)
	route.HandleFunc("/admin/category/new", model.NewCategory)
	route.HandleFunc("/admin/category/edit", model.EditCategory)
	route.HandleFunc("/admin/category/insert", model.InsertCategory)
	route.HandleFunc("/admin/category/delete", model.DeleteCategory)

	route.HandleFunc("/keikibook", models.HomeHome)
	route.HandleFunc("/keikibook/login", models.ShowLogin)
	route.HandleFunc("/keikibook/signup", models.ShowSignUp)
	route.HandleFunc("/keikibook/signups", models.Signup)
	route.HandleFunc("/keikibook/logins", models.Signin)
	route.HandleFunc("/keikibook/bug", models.Bug)
	route.HandleFunc("/keikibook/logout", models.LogoutHandler)

	log.Fatal(http.ListenAndServe(":8080", route))
}
