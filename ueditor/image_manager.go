package ueditor

import (
	"path"
	"net/http"
)

func listImage( response http.ResponseWriter, request *http.Request  )  {
	var picPathList,start, total = getList(request,path.Join("static", "upload","image"))
	listResult( picPathList,start, total, response)
}