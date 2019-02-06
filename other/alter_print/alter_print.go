package main

import (
	"fmt"
	"sync"
)

func main() {
	var chanN chan struct{} = make(chan struct{})
	var chanC chan struct{} = make(chan struct{})
	//var done chan struct{} = make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i += 2 {
			fmt.Print(i)
			fmt.Print(i + 1)
			chanN <- struct{}{}
			<-chanC
		}
	}()
	go func() {
		defer wg.Done()
		number := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		for i := 0; i < 10; i += 2 {
			<-chanN
			fmt.Print(number[i])
			fmt.Print(number[i+1])
			chanC <- struct{}{}
		}
		//done <- struct{}{}
	}()
	//<-done
	wg.Wait()
}
