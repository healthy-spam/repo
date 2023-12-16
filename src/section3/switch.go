package main

import "fmt"

func main() {
	//제어문(조건문) - switch
	//Switch 뒤 표현식(Expression) 생략가능
	//case 뒤 표현식(Expression) 사용 가능
	//자동 break 때문에 fallthrough whswo
	//Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

	//예제1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	//예제2
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang!")
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("None")
	}

	//예제3
	switch i, j := 10, 20; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j와 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}
}
