package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Go-routines")
	go func() {
		defer wg.Done()
		// Displays alphabets (lowercase) 3 times
		for count:=0; count< 3; count++ {
			for char:= 'a';char<'a'+26;char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	go func() {
		defer wg.Done()
		// Displays alphabets (uppercase) 3 times
		for count:=0; count< 3; count++ {
			for char:= 'A';char<'A'+26;char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating program")
}