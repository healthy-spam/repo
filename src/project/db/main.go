package main

import (
	"fmt"
	"net/http"
)

func a(writer http.ResponseWriter, request *http.Request) {

	_, err := writer.Write([]byte("Hello World!"))
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(request.URL)
}

func handler() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", a)

	return mux

}

func main() {

	addr := "localhost:8080"

	err := http.ListenAndServe(addr, handler())
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
