package admin

import (
	"log"
	"net/http"

	db "github.com/JIeeiroSst/config"
	upload "github.com/JIeeiroSst/utils"
)

type Category struct {
	Id   int
	Name string
}

func IndexCategory(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDb, err := db.Query("select * from category")
	if err != nil {
		panic(err.Error())
	}
	cat := Category{}
	res := []Category{}
	for selDb.Next() {
		var id int
		var name string
		err = selDb.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		cat.Id = id
		cat.Name = name
		res = append(res, cat)
	}
	log.Println(res)
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "CategoryIndex", res)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func ShowCattegory(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDb, err := db.Query("select * from category where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	cat := Category{}
	for selDb.Next() {
		var id int
		var name string
		err = selDb.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		cat.Id = id
		cat.Name = name
	}
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "CategoryShow", cat)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func NewCategory(w http.ResponseWriter, r *http.Request) {
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "CategoryNew", nil)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDb, err := db.Query("select * from category where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	cat := Category{}
	for selDb.Next() {
		var id int
		var name string
		err = selDb.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		cat.Id = id
		cat.Name = name
	}
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "CategoryEdit", cat)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		insert, err := db.Prepare("insert into category(name) values(?)")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(name)
		log.Println("insert category" + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin/category", 301)
}

func UPdateCategory(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		id := r.FormValue("id")
		update, err := db.Prepare("upadte Category set name=? where id=?")
		if err != nil {
			panic(err.Error())
		}
		update.Exec(name, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin/category", 301)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	delete, err := db.Prepare("delete from category where id=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(nId)
	defer db.Close()
	http.Redirect(w, r, "/admin/category", 301)
}
