package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello,world\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", helloHandler)
	http.ListenAndServe(":3000", mux)
}
