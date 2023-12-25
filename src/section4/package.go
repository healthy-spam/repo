package main

import (
	"fmt"
	"os"
)

func main() {

	var name string

	fmt.Println("이름은? :")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
