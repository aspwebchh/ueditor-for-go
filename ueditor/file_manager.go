package ueditor

import (
	"net/http"
)

func listFile( response http.ResponseWriter, request *http.Request )  {
	var picPathList,start, total = getList(request, config.FilePath)
	listResult( picPathList,start, total, response)
}