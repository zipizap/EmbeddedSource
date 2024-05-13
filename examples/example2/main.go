package main

import (
	"net/http"
)

func main() {
	embeddedSourceStub()

	port := ":9999"
	handler := http.FileServer(http.Dir('.'))
	http.ListenAndServe(port, handler)
}
