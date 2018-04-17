package ueditor

import (
	"path"
	"os"
	"io"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strconv"
)

func newFileName() string  {
	var now = time.Now()
	var fileName = now.Format("200601021504")
	fileName = fileName + strconv.Itoa( now.Nanosecond() )
	return fileName
}

func upload(pathString string, response http.ResponseWriter, request *http.Request) {
	var file, header, err = request.FormFile("upfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var filename = newFileName() +  path.Ext(header.Filename)
	err = os.MkdirAll(pathString, 0775)
	if err != nil {
		panic(err)
	}
	targetFile, err := os.Create(path.Join(pathString, filename))
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	io.Copy(targetFile, file)
	var responseContent = map[string]string{
		"url":      fmt.Sprintf("/"+ pathString +"/%s", filename),
		"title":    "",
		"original": header.Filename,
		"state":    "SUCCESS",
	}
	var  responseJSON, _ = json.Marshal(responseContent)
	response.Write(responseJSON)
}
