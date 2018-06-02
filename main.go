package main

import 	"fmt"

func main() {
	fmt.Println("Hello World")
	hello("passing param")
	fmt.Println(helloReturn("test"))
	a, b := x(1, 2)
	fmt.Println(a, b)

	inner := outer("Thanakorn")
	innerString := inner()
	fmt.Println(innerString)
}

func hello(s string) {
	fmt.Println("void call func hello", s)
}

func helloReturn(s string) string {
	return fmt.Sprintf("return call func helloReturn %s", s)
}

func x (a int, b int) (int, int) {
	return a, a+b
}

func outer (name string) func() string {
	
	text := "Hello" + name

	return func() string {
		return text
	}
}
