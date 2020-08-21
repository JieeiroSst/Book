package utils

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request, name string) string {
	file, handle, err := r.FormFile(name)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		saveFile(w, r, file, handle)
	case "image/png":
		saveFile(w, r, file, handle)
	case "application/pdf":
		saveFile(w, r, file, handle)
	case "image/gif":
		saveFile(w, r, file, handle)
	default:
		jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
	}
	return handle.Filename
}

func saveFile(w http.ResponseWriter, r *http.Request, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = ioutil.WriteFile("./uploads/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	http.Redirect(w, r, "/admin/book", 301)
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
