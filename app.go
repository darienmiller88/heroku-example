package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":5000", nil)

}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>server is up</h1>")
}
