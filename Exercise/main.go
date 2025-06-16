// interview question
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // / multiple goroutines
// // PUSH DATA
// // PRINT DATA
// func main() {
// 	datachan := make(chan string, 10)

// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			for j := 0; j <= 3; j++ {
// 				messaage := fmt.Sprint("Message from Microservice  ", id, " - ", j)
// 				datachan <- messaage
// 			}
// 		}(i)
// 	}

// 	// Consumer goroutine
// 	go func() {
// 		for msg := range datachan {
// 			fmt.Println("Hi Data Consumer: ", msg)
// 		}
// 	}()

// 	wg.Wait()
// 	close(datachan)
// }

///---------------------------------------------------------------------------------

// RACE CONDITION

// package main

// import (
// 	"fmt"
// 	"time"
// 	"mu"
// )

// var counter int

// func increment() {
// 	for i := 0; i < 1000; i++ {
// 		mu.Lock() // Lock the mutex before accessing the counter
// 		counter++ // Not safe for concurrent access
// 		mu.Unlock() // Unlock the mutex after accessing the counter
// 	}
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go increment()
// 	}
// 	time.Sleep(1 * time.Second) // Let goroutines finish
// 	fmt.Println("Final counter (with race condition):", counter)
// }

// -----------------------------------------------------------------------------------

// RACE CONDITION

// package main

// import (
// 	"fmt"
// 	"time"
// )

// var counter = 0

// func increment() {
//     counter++
// }

// func main() {
//     for i := 0; i < 1000; i++ {
//         go increment()  // Multiple goroutines trying to increment the counter
//     }

//		time.Sleep(3*time.Second) // Wait for goroutines to finish
//	    fmt.Println("Final counter:", counter)  // This may not give correct result due to race condition
//	}
//

///---------------------------------------------------------------------------------

// Buffered VS UnBuffered Channels In Golang  - Video

// package main
// import "fmt"

// func main() {

// 	msgch := make(chan int,10) //buffered
// 	for i:=20 ; i<29; i++ { //sending 10 values	to unbuffered channel
// 		msgch <- i
// 		fmt.Println("msgch: ", i  )
// 	}
// 	// msgch<-34   //sending 1 value to buffered channel
// 	// go func() {
//     //     msgch <- 34            // Send value 34 to the channel in a goroutine
//     // }()
// 	output := <-msgch
// 	fmt.Println("Received msgch: ", output)
// }

//In Go, when working with goroutines and channels, you'll often hear the terms:
// 🔵 Unbuffered channels (Synchronous	)
// 🟢 Buffered channels (Asynchronous (up to buffer capacity )
// These terms define how data is sent & received across goroutines using channels.
// Unbuffered channel has no capacity.
// A send channel (chan<-) blocks until another goroutine is ready to receive (chan<-) the value and vice versa.
// THis means that both sender and receiver must be ready at same time --- synchronous communication.

// package main

// import "fmt"

// func main() {
// 	ch := make(chan string, 3) // buffered with capacity 2

// 	ch <- "World1"
// 	ch <- "World2"
// 	ch <- "Hello3"

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// No goroutines needed here — since the channel can hold 3 values, the sends don’t block immediately.

// /---------------------------------------------------------------------------------

// // Goroutines , Waitgroup
// // Goroutines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent tasks in Go programs.

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func sayHi() {
// 	fmt.Println("Hi, Sushant!")
// 	time.Sleep(1 * time.Second) // Simulate some work
// 	fmt.Println("Bye, Sushant!")
// }

// func sayHello() {
// 	fmt.Println("Hello, World!")
// 	time.Sleep(2 * time.Second) // Simulate some work
// 	fmt.Println("Hello, World Ended!")
// }

// func task(id int, w *sync.WaitGroup) {
// 	defer w.Done()
// 	fmt.Println("Worker", id)
// }

// func main() {
// 	fmt.Println("Learning Golang...")

// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go task(i, &wg) //blocking
// 	}

// 	go sayHello()
// 	go sayHi()

// 	wg.Wait()
// 	// time.Sleep(1 * time.Millisecond) // Wait for goroutine to finish
// }

// /---------------------------------------------------------------------------------

// // Channels in Golang

// package main

// import (
// 	"fmt"
// )

// // receiving
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// // sending
// func processNum(numChan chan int) {
// 	for {
// 		num := <-numChan
// 		fmt.Println("Received number:", num)
// 	}
// }

// func main() {

// 	//Receiving
// 	result := make(chan int)
// 	go sum(result, 4, 5)
// 	res := <-result //blocking
// 	fmt.Println("Sum:", res)

// 	// Sending
// 	numChan := make(chan int)
// 	go processNum(numChan)
// 	for i := 20; i <= 30; i++ {
// 		numChan <- i
// 	}
// 	close(numChan)

// }

// /---------------------------------------------------------------------------------

// - WaitGroup Example :
//  Waitgroup is a mechnism to synchronize the completion of multiple goroutines in Go.
//  WaitGroup is used to wait for a collection of goroutines to finish executing.
// - It provides a way to synchronize the completion of multiple goroutines.
// - It is part of the sync package in Go and is commonly used in concurrent programming to
// ensure that the main program does not exit before all spawned goroutines have finished their execution.
// - It provides a way to block the main goroutine (or any other goroutine) until all the goroutines it is waiting for have completed their tasks.

// // WaitGroup is a synchronization primitive that allows you to wait for a collection of goroutines to finish executing.
// // It is part of the sync package in Go and is commonly used in concurrent programming to ensure that the main program does not exit before all spawned goroutines have finished

// 1. Add(int) - Increment the WaitGroup counter by one. This is usually called before starting a new goroutine to indicate that there is a new task to wait for.
// 2. Done() - Decrement the WaitGroup counter by one. This is usually called at the end of a goroutine to signal that the goroutine has completed its task
// 3. Wait() - Block the execution of main(or calling) goroutine until the WaitGroup counter is zero. This means that all the goroutines that were added to the WaitGroup have completed their tasks and the program can safely exit. This is typically called in the main goroutine to wait for all other goroutines to finish.

// - Create a sync.WaitGroup variable (e.g., var wg sync.WaitGroup) to initialize the WaitGroup.
// -** Always call wg.Add(1) before starting a goroutine to avoid race conditions.
// - Always use defer wg.Done() at the beginning of the goroutine to ensure the counter is decremented, even if the function exits early due to an error.

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func sayHi(wg *sync.WaitGroup) {
// 	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
// 	fmt.Println("Hi, Sushant!")
// 	time.Sleep(1 * time.Second) // Simulate some work
// 	fmt.Println("Bye, Sushant!")
// }
// func sayHello(wg *sync.WaitGroup) {
// 	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
// 	fmt.Println("Hello, World!")
// 	// time.Sleep(2 * time.Second) // Simulate some work
// 	fmt.Println("Hello, World Ended!")
// }
// func main() {
// 	fmt.Println("Learning Golang...")
// 	var wg sync.WaitGroup // Create a WaitGroup to wait for goroutines to finish
// 	wg.Add(1)
// 	go sayHello(&wg)

// 	wg.Add(1) // Increment the WaitGroup counter for the second goroutine
// 	go sayHi(&wg)

// 	time.Sleep(1 * time.Millisecond) // Wait for goroutine to finish

// 	wg.Wait() // Wait for all goroutines to finish
// 	fmt.Println("All goroutines completed...")
// }

/// ----------------------------------------------------------------------------------

//// Interfaces in Golang

// package main

// import "fmt"

// //define an interface
// // interface is a collection of method signatures
// // it defines a contract that a type must fulfill to implement the interface
// // an interface is a type that specifies a contract that other types must implement

// // SHAPE INTERFACE
// type Shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type Circle struct {
// 	Radius float64
// }
// type Rectangle struct {
// 	Length, Width float64
// }
// type Square struct {
// 	Side float64
// }

// // Implement Shape interface for Circle
// func (c Circle) area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func (c Circle) perimeter() float64 {
// 	return 2 * 3.14 * c.Radius
// }

// // Implement Shape interface for Rectangle
// func (r Rectangle) area() float64 {
// 	return r.Length * r.Width
// }
// func (r Rectangle) perimeter() float64 {
// 	return 2 * (r.Length + r.Width)
// }

// // Implement Shape interface for Square
// func (s Square) area() float64 {
// 	return s.Side * s.Side
// }
// func (s Square) perimeter() float64 {
// 	return 4 * s.Side
// }

// // INTERFACE
// func printShapeDetails(s Shape) {
// 	fmt.Println("Area:", s.area())
// 	fmt.Println("Perimeter:", s.perimeter())
// 	fmt.Println("-------------------------")
// }

// func main() {
// 	circle := Circle{Radius: 5}
// 	rectange := Rectangle{Length: 4, Width: 6}

// 	// fmt.Println("Circle:", circle)
// 	printShapeDetails(circle)
// 	printShapeDetails(rectange)

// 	// fmt.Println("Rectangle:", rectange)
// }

// -----------------------------------------------------------------------------------

///  MCQS

// package main

// import "fmt"

// func main() {
// 	evenCh := make(chan int)
// 	oddCh := make(chan int)

// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)

// 	for i := 0; i <= 10; i++ {
// 		if i%2 == 0 {
// 			evenCh <- i
// 		} else {
// 			oddCh <- i
// 		}
// 	}

// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// 	go print(oddCh, evenCh)
// }

// func print(oddCh, evenCh <-chan int) {
// 	fmt.Println("Printing odd and even numbers...")
// 	for {
// 		select {

// 		case i := <-oddCh:
// 			fmt.Println(i)

// 		case i := <-evenCh:
// 			fmt.Println(i)
// 		}
// 	}
// }

///-------------------------

// CHATGPT:

// 1 Goroutine
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("(sayHello) Hello from goroutine 1")
}

func sendMessage(ch chan string) {
	fmt.Println("(sendMessage) Sending message to channel ...")
	ch <- "Hello from goroutine"
	fmt.Println("(sendMessage) Message sent to channel ...")
}

func worker(id int, ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("Worker %d: Task %d completed", id, i+1)
	}
}

func main() {
	fmt.Println("------ goroutine ------")
	go sayHello()
	fmt.Println("(main) Main function is running")
	time.Sleep(1 * time.Second)
	fmt.Println("(main) Main function is exiting")

	fmt.Println("------ channel ------")
	messageChannel := make(chan string)
	go sendMessage(messageChannel)
	newmessage := <-messageChannel
	fmt.Println("(main) Received message:", newmessage)

	fmt.Println("------ ✅ Example: Multiple Goroutines Communicating via Channels ------")

	// Multiple goroutines, single channel
	ch := make(chan string)

	// Start 2 workers
	go worker(1, ch)
	go worker(2, ch)

	// Read 6 messages (2 workers × 3 messages each)
	for i := 0; i < 6; i++ {
		msg := <-ch
		fmt.Println("(main) Received:", msg)
	}
}

// Answer:
// Yes, you can use channels instead of sync.WaitGroup to wait for goroutines to finish — but that doesn’t mean you always should.
