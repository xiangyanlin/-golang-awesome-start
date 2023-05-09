package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 在/后面加上 index ，来指定访问路径
	http.HandleFunc("/index", index)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}
