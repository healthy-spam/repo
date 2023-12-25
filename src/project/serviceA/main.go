package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {

	resp, err := http.Get("http://localhost:8081/")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	writer.Write(bodyByte)

	fmt.Println("URL: ", request.URL)
}

func handler() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", mainHandler)

	return mux

}

func main() {

	addr := "localhost:8082"

	err := http.ListenAndServe(addr, handler())
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

}
