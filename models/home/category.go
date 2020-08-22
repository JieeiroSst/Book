package home

import (
	"log"
	"net/http"

	db "github.com/JIeeiroSst/config"
	cat "github.com/JIeeiroSst/models/admin"
)

func MenuCategory(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	seldb, err := db.Query("select * from category")
	if err != nil {
		panic(err.Error())
	}
	cats := cat.Category{}
	res := []cat.Category{}
	for seldb.Next() {
		var id int
		var name string
		err = seldb.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		cats.Id = id
		cats.Name = name

		res = append(res, cats)
	}
	log.Println(res)
	tmpl.ExecuteTemplate(w, "HomeMenu", res)
	defer db.Close()
}
