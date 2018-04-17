package ueditor

import (
	"path"
	"net/http"
)

func listFile( response http.ResponseWriter, request *http.Request )  {
	var picPathList,start, total = getList(request, path.Join("static", "upload","file"))
	listResult( picPathList,start, total, response)
}