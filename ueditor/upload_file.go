package ueditor

import (
	"net/http"
	"path"
)

func uploadFile(response http.ResponseWriter, request *http.Request) {
	upload(path.Join("static", "upload","file"),response,request)
}

