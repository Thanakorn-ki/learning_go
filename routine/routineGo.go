package main

import (
	"fmt"
	"sync"
)

func main() {
	var w = sync.WaitGroup{}
	w.Add(2)
	go func() {
		fmt.Println("routine1")
		w.Done()
	}()
	go func() {
		fmt.Println("routine2")
		w.Done()
	}()
	fmt.Println("hi")
	w.Wait()
}
