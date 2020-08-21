package admin

import (
	"net/http"

	db "github.com/JIeeiroSst/config"
	upload "github.com/JIeeiroSst/utils"
)

type User struct {
	Id       int
	Name     string
	Password string
	EmpId    int
}

func IndexUser(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDb, err := db.Query("select * from User")
	if err != nil {
		panic(err.Error())
	}

	user := User{}
	res := []User{}

	for selDb.Next() {
		var id, empId int
		var name, password string
		err = selDb.Scan(&id, &name, &password, &empId)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Name = name
		user.Password = password
		user.EmpId = empId

		res = append(res, user)
	}
	userName := upload.GetUserName(r)
	if userName != "" {
		tmpl.ExecuteTemplate(w, "UserShow", res)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
	defer db.Close()
}
