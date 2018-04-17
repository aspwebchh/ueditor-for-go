package main

import (
	"net/http"
	"./ueditor"
	)
func main()  {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ueditor/go", ueditor.UEditor)
	http.ListenAndServe(":10079" , nil)
}
