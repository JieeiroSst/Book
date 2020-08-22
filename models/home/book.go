package home

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

func IndexBookHome(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("select * from Book")
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
		tmpl.ExecuteTemplate(w, "HomeMain", res)
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func BookDetail(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDb, err := db.Query("select id,file from Book where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	book := Book{}
	for selDb.Next() {
		var id int
		var file string
		err = selDb.Scan(&id, &file)
		log.Println(file)
		if err != nil {
			panic(err.Error())
		}
		book.Id = id
		book.File = file
	}
	log.Println(book.File)
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "HomeBook", book)
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}

func BookCat(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select k.id,k.name,k.authId,k.file,k.content from category c join book k on c.id=k.authId where c.id=?", nId)
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
		tmpl.ExecuteTemplate(w, "HomeCat", res)
	} else {
		http.Redirect(w, r, "/keikibook/login", 302)
	}
	defer db.Close()
}
