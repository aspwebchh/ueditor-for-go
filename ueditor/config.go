package ueditor

import (
	"bytes"
	"os"
	"net/http"
)

func config(response http.ResponseWriter, request *http.Request) {
	var file, err = os.Open("config_ue.json")
	var configJson  []byte
	if err != nil {
		configJson = []byte("{}")
	} else {
		defer file.Close()
		buf := bytes.NewBuffer(nil)
		buf.ReadFrom(file)
		configJson = buf.Bytes()
	}
	response.Write(configJson)
}