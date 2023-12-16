package main

import "fmt"

func main() {
	//제어문(조건문)
	//IF 문 : 반드시 BooLean 검사 -> 1, 0(사용불가 : 자동 형변환 불가)

	var a int = 10
	b := 10

	if a > 8 {
		fmt.Println("9이상")
	}

	if b >= 13 {
		fmt.Println("14이상")
	}

	//에러 발생1
	/*if b >= 13
	{
		fmt.Println("14이상")
	}*/

	//에러 발생2
	/*if b >= 13
	fmt.Println("14이상")*/

	//에러 발생3
	/*if c := 1; c {
		fmt.Println("true")
	}*/
}
