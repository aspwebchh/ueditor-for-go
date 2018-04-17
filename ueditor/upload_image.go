package ueditor

import (
	"net/http"
	"path"
)

func uploadImage(response http.ResponseWriter, request *http.Request) {
	upload(path.Join("static", "upload","image"),response,request)
}
