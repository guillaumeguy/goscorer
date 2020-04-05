package main

import (
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello there")
}

func main() {
	var port = os.Getenv("PORT")
	println(port)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":" + port, nil)
}
