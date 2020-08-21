package admin

import (
	"log"
	"net/http"

	db "github.com/JIeeiroSst/config"
	upload "github.com/JIeeiroSst/utils"
)

type Book struct {
	Id      int
	Name    string
	AuthId  int
	File    string
	Content string
}

func IndexBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("select * from Book order by id desc")
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	res := []Book{}

	for selDB.Next() {
		var id, authId int
		var name, file, content string
		err = selDB.Scan(&id, &name, &authId, &file, &content)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.Name = name
		book.AuthId = authId
		book.File = file
		book.Content = content

		res = append(res, book)
	}
	log.Println(res)
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "BookIndex", res)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func ShowBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from Book where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	for selDB.Next() {
		var id, authId int
		var name, file, content string
		err = selDB.Scan(&id, &name, &authId, &file, &content)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.Name = name
		book.AuthId = authId
		book.File = file
		book.Content = content
	}
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "BookShow", book)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func NewBook(w http.ResponseWriter, r *http.Request) {
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "BookNew", nil)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from Book where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	for selDB.Next() {
		var id, authId int
		var name, file, content string
		err = selDB.Scan(&id, &name, &authId, &file, &content)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.Name = name
		book.AuthId = authId
		book.File = file
		book.Content = content
	}
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "BookEdit", book)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}

func InsertBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		authId := r.FormValue("authId")
		file := upload.UploadFile(w, r, "file")
		content := r.FormValue("content")
		insForm, err := db.Prepare("insert into Book(name,authId,file,content)values(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, authId, file, content)
		log.Println("Insert Name:" + name + "|authID:" + authId + "|content:" + content)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin/book", 301)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		authId := r.FormValue("authId")
		file := upload.UploadFile(w, r, "file")
		content := r.FormValue("content")
		id := r.FormValue("id")
		insForm, err := db.Prepare("update Book set name=?,authId=?,file=?,content=? where id=?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, authId, file, content, id)
		log.Println("Insert Name:" + name + "|authID:" + authId + "|file:" + content)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin/book", 301)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	delForm, err := db.Prepare("delete from Book where id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nId)
	log.Println("delete")
	defer db.Close()
	http.Redirect(w, r, "/admin/book", 301)
}
