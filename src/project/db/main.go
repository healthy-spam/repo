package main

import (
	"fmt"
	"net/http"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {

	_, err := writer.Write([]byte("I'm data"))
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(request.URL)
}

func handler() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", mainHandler)

	return mux

}

func main() {

	addr := "localhost:8080"

	err := http.ListenAndServe(addr, handler())
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
