package main

import "fmt"

func goRoutines() {

	var myChannel = make(chan int)
	//fmt.Printf("Printed value is; %T ", myChannel)
	//fmt.Printf("Printed value is; %v ", myChannel)

	go calculateSquares(myChannel)

	/*for {
		val, ok := <-myChannel
		if ok == false {
			fmt.Println(val, ok, "loop broke")
			break
		} else {
			fmt.Println(val, ok)
		}
	}*/

	for val := range myChannel {
		fmt.Println(val)
	}

	fmt.Println("Main go routine is stopped")

}

func printSquaresResults(computationChannel chan int) {
	for {
		val, ok := <-computationChannel
		if ok == false {
			fmt.Println(val, ok, "loop broke")
			break
		} else {
			fmt.Println(val, ok)
		}
	}
}

func calculateSquares(computationChannel chan int) {
	for i := 0; i <= 9; i++ {
		computationChannel <- i * i
	}
	close(computationChannel)
}

func greet(greetingChannel chan string) {
	fmt.Printf("Greeetings!!! " + <-greetingChannel)

	<-greetingChannel
}
