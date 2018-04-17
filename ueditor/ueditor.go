package ueditor

import (
	"net/http"
)

const ACTION_CONFIG = "config"
const ACTION_LIST_IMAGE = "listimage"
const ACTION_LIST_FILE = "listfile"
const ACTION_UPLOAD_IMAGE = "uploadimage"
const ACTION_UPLOAD_FILE = "uploadfile"

func UEditor(response http.ResponseWriter, request *http.Request) {
	var action = request.FormValue("action")
	if action == ACTION_CONFIG {
		config(response, request)
	} else if action == ACTION_LIST_IMAGE {
		listImage(response, request)
	} else if action == ACTION_LIST_FILE {
		listFile(response, request)
	} else if action == ACTION_UPLOAD_IMAGE {
		uploadImage(response, request)
	} else if action == ACTION_UPLOAD_FILE {
		uploadFile(response, request)
	}
}
