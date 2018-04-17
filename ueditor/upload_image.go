package ueditor

import (
	"net/http"
)

func uploadImage(response http.ResponseWriter, request *http.Request) {
	upload(config.ImagePath,response,request)
}
