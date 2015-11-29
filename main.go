package main

import (
	"io"
	"net/http"

	"github.com/jd1123/adproxy/stopper"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {

	http.HandleFunc("/", hello)
	stopper.StoppableListenAndServe(":9999", nil)
}
