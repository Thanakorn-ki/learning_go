package main

import "fmt"

type justSignal struct {
	name string
}

func main() {
	ch1 := make(chan justSignal)
	ch2 := make(chan justSignal)
	go justPrint(ch1, "test1")
	fmt.Printf("%#v", <-ch1)
	go justPrint(ch2, "test1")
	chValue := <-ch2
	fmt.Printf("%#v", chValue.name)
}

func justPrint(cch chan justSignal, name string) {
	cch <- justSignal{name: name}
}
