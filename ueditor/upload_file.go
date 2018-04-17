package ueditor

import (
	"net/http"
)

func uploadFile(response http.ResponseWriter, request *http.Request) {
	upload(config.FilePath,response,request)
}

