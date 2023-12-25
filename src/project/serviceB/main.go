package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	// 메인 핸들러 로직
	fmt.Fprint(writer, "Hello, World!")
}

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	// 헬스체크 핸들러 로직
	fmt.Fprint(writer, "Health check passed")
}

func healthCheck() {
	// 5초 주기의 ticker 생성
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 서버가 살아있음을 5초마다 콘솔에 출력
			fmt.Println("서버가 살아있습니다.")

			// 헬스체크 핸들러 호출 (실제 헬스체크를 수행하는 부분)
			client := &http.Client{
				Timeout: 5 * time.Second,
			}

			resp, err := client.Get("http://localhost:8081/health")
			if err != nil {
				fmt.Println("헬스체크 실패:", err)
			} else {
				defer resp.Body.Close()
				fmt.Println("헬스체크 결과:", resp.Status)
			}
		}
	}
}

func handler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/health", healthCheckHandler)
	return mux
}

func main() {
	addr := "localhost:8081"

	// 헬스체크를 주기적으로 수행하는 고루틴 시작
	go healthCheck()

	err := http.ListenAndServe(addr, handler())
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
