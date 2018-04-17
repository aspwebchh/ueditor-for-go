package ueditor

import (
	"net/http"
)

func listImage( response http.ResponseWriter, request *http.Request  )  {
	var picPathList,start, total = getList(request, config.ImagePath)
	listResult( picPathList,start, total, response)
}