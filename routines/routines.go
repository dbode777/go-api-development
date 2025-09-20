package main

import (
	"fmt"
	"sync"
	"time"
)

func run1() {
	time.Sleep(2 * time.Second)
	fmt.Println("Routine 1")
}

func run2() {
	time.Sleep(4 * time.Second)
	fmt.Println("Routine 2")
}

func run3() {
	time.Sleep(6 * time.Second)
	fmt.Println("Routine 3")
}

// Routines can wait for an execution if a channel is passed to them
func add(x int, y int, ch chan int) {
	ch <- x + y
}

// Channels can have a specified direction whether it can only send or recieve values
func printStatment(ch <-chan string) string {
	x := <-ch // arrow on the left side indicates recieving
	return x
}

// If nothing is sent or recieved to a channel, it will error out with a deadlock error.
// To avoid this, use a buffered channel (see ch4)
func sublist(list []int, ch chan<- []int, start int, end int) {
	ch <- list[start:end] // arrow on the right side indicates sending
}

type Counter struct {
	value int
	lock  sync.Mutex // Locks are used to halt thresholds from executing on the same memory address and wait for other threads to finish their executions
}

func increment(c *Counter, ch chan<- bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
	fmt.Println(c.value)
	ch <- true // This is one way to send a value without having to use the time.Sleep function
}

// Routines using 'go' are executed concurrently. Interior functions won't complete if the exterior function completes first.
func main() {
	startTime := time.Now()

	fmt.Println("Start time: ", startTime)
	ch := make(chan int)
	ch2 := make(chan int)

	go run1()
	go run2()
	go run3() // if you remove go, each of these functions will be executed from the top down fully.
	go add(1, 2, ch)
	x := <-ch // Wait for the channel to return a value

	go add(1, 5, ch2)

	// Use select statements to wait for a value to be returned from at least one of multiple channels
	// These values will be returned whenever one of the channels finishes executing
	select {
	case x = <-ch:
		fmt.Println(x)
	case y := <-ch2:
		fmt.Println(y)
	}

	ch3 := make(chan string, 2) // Must include buffer size here to avoid deadlock error on line 67
	ch3 <- "Hello"
	fmt.Println(printStatment(ch3))

	ch4 := make(chan []int, 2) // Can now store up to 2 values in the channel
	go sublist([]int{1, 2, 3, 4, 5}, ch4, 1, 3)
	a := <-ch4

	ch4 <- []int{6, 7, 8, 9, 10}
	b := <-ch4

	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)

	counter := Counter{value: 0}
	ch5 := make(chan bool) // can also use sync.WaitGroup & sync.Add(number of channels needed for operations) to avoid using time.Sleep
	for i := 0; i < 10; i++ {
		go increment(&counter, ch5)
	}
	for i := 0; i < 10; i++ {
		<-ch5
	}

	time.Sleep(7 * time.Second) // Adjust to a number greater than or equal to 6 to see the result of run3()

	fmt.Println("Done. Duration of execution: ", time.Since(startTime))
}
