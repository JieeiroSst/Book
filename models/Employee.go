package models

import (
	"html/template"
	"log"
	"net/http"

	db "github.com/JIeeiroSst/config"
	session "github.com/JIeeiroSst/utils"
)

type Employee struct {
	Id   int
	Name string
	City string
}

var tmpl = template.Must(template.ParseGlob("views/admin/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	userName := session.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "EmployeeIndex", res)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	userName := session.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "EmployeeShow", emp)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	userName := session.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "EmployeeNew", nil)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	userName := session.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "EmployeeEdit", emp)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin/employe", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin/employe", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/admin/employe", 301)
}
