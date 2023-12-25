package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	_, err = writer.Write(byteBody)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func healthCheck() {

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			resp, err := http.Get("http://localhost:8082")
			if err != nil {
				fmt.Println("헬스체크 실패: ", err)
			} else {
				fmt.Println("헬스체크 성공: ", resp.StatusCode)
			}
		}
	}
}

func handler() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", mainHandler)

	return mux
}

func main() {

	addr := "localhost:8081"

	go healthCheck()

	err := http.ListenAndServe(addr, handler())
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
