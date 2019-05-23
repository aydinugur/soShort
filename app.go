package main

import "net/http"

func sayHello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello world!"))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8080", nil)
}
