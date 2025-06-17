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

// GPT:

// // 1 Goroutine
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHello() {
// 	fmt.Println("(sayHello) Hello from goroutine 1")
// }

// func sendMessage(ch chan string) {
// 	fmt.Println("(sendMessage) Sending message to channel ...")
// 	ch <- "Hello from goroutine"
// 	fmt.Println("(sendMessage) Message sent to channel ...")
// }

// func worker(id int, ch chan string) {
// 	for i := 0; i < 3; i++ {
// 		ch <- fmt.Sprintf("Worker %d: Task %d completed", id, i+1)
// 	}
// }

// func main() {
// 	fmt.Println("------ goroutine ------")
// 	go sayHello()
// 	fmt.Println("(main) Main function is running")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("(main) Main function is exiting")

// 	fmt.Println("------ channel ------")
// 	messageChannel := make(chan string)
// 	go sendMessage(messageChannel)
// 	newmessage := <-messageChannel
// 	fmt.Println("(main) Received message:", newmessage)

// 	fmt.Println("------ ✅ Example: Multiple Goroutines Communicating via Channels ------")

// 	// Multiple goroutines, single channel
// 	ch := make(chan string)

// 	// Start 2 workers
// 	go worker(1, ch)
// 	go worker(2, ch)

// 	// Read 6 messages (2 workers × 3 messages each)
// 	for i := 0; i < 6; i++ {
// 		msg := <-ch
// 		fmt.Println("(main) Received:", msg)
// 	}
// }

// // Answer:
// // Yes, you can use channels instead of sync.WaitGroup to wait for goroutines to finish — but that doesn’t mean you always should

///--------------------------------------------------

// PRACTICE QUESTIONS


1)Even Odd using 2 Channel & 1 Goroutine

package main
import "fmt"

func main() {
    evenCh := make(chan int)
    oddCh := make(chan int)
    
    go print(oddCh,evenCh)
    for i:=0; i<=10; i++ {
        if i%2==0 {
            evenCh<-i
        }else{
            oddCh<-i
        }
    }
}

func print(oddCh, evenCh <-chan int){
    for {
        select {
            case i:=<-oddCh :
            fmt.Println(i)
            case i:=<-evenCh :
            fmt.Println(i)
        }
    }
}



















2)Even Odd using 2 Channel & 2  Goroutine

package main

import (
	"fmt"
	"sync"
)

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go even(evenCh, &wg)
	go odd(oddCh, &wg)
	go func() {
		wg.Wait()
		close(evenCh)
		close(oddCh)
	}()

	for i := 0; i < 6; i++ {
		fmt.Println(<-evenCh)
		fmt.Println(<-oddCh)
	}

}

func even(evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i = i + 2 {
		evenCh <- i
	}
}

func odd(oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i = i + 2 {
		oddCh <- i
	}
}


3)Even Odd using 1 Channel & 2  Goroutine


package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)
	go Even(ch)
	ch <- true
	go Odd(ch)

	time.Sleep(time.Second * 10)

}

func Odd(c chan bool) {
	for i := 1; i < 10; {
		<-c
		fmt.Println("value is Odd ", i)
		i += 2
		c <- true
	}
}

func Even(c chan bool) {
	for i := 0; i < 10; {
		<-c
		fmt.Println("value is even ", i)
		i += 2
		c <- true
	}
}









4)Even Odd using 1 Channel & 2  Goroutine

package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)
	go Even(ch)
	ch <- true
	go Odd(ch)

	time.Sleep(time.Second * 10)

}

func Odd(c chan bool) {
	for i := 1; i < 10; {
		<-c
		fmt.Println("value is Odd ", i)
		i += 2
		c <- true
	}
}

func Even(c chan bool) {
	for i := 0; i < 10; {
		<-c
		fmt.Println("value is even ", i)
		i += 2
		c <- true
	}
}










5)Semaphore 

func worker(id int, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}
	fmt.Printf("Worker %d started\n", id)

	fmt.Printf("Worker %d finished\n", id)
	<-sem // Release semaphore (increase resource count)
}

func main() {
	sem := make(chan struct{}, 3)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, sem)
	}
	wg.Wait()
}


6)Worker Pool 
func worker(id int, jobs <-chan int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
 }
7 ) Batch 

package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	batch := 5
	var wg sync.WaitGroup

	// Create a buffered channel to store the sums of each batch
	ch := make(chan int, len(data)/batch+1)

	for i := 0; i < len(data); i += batch {
		end := i + batch
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go add(data[i:end], ch, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(ch)

	// Collect and sum up the results from the channel
	totalSum := 0
	for sum := range ch {
		totalSum += sum
	}

	fmt.Println("Total : “ totalSum)
}

func add(arr []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, n := range arr {
		sum += n
	}
	ch <- sum
}






8)Increment Atomic Package 

func main() {
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}


























9)Select 


func main() {
	ch1 := make(chan []int)
	ch2 := make(chan []int)
	ch3 := make(chan []int)

	go imei1(ch1)
	go imei2(ch2)
	go imei3(ch3)

	t := time.After(4 * time.Second)

	for {
		select {
		case i := <-ch1:
			fmt.Println("IMEI 1 :", i)
		case i := <-ch2:
			fmt.Println("IMEI 2 :", i)
		case i := <-ch3:
			fmt.Println("IMEI 3 :", i)
		case <-t:
			fmt.Println("TIMEOUT")
			return
		}

	}
}

func imei1(ch chan []int) {
	time.Sleep(1 * time.Second)
	s := []int{1, 2, 3, 4}
	ch <- s
}

func imei2(ch chan []int) {
	time.Sleep(2 * time.Second)
	s := []int{5, 6, 7, 8}
	ch <- s
}

func imei3(ch chan []int) {
	time.Sleep(5 * time.Second)
	s := []int{9, 10, 11, 12}
	ch <- s
}
10)Producer & Consumer 

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		producerWg.Add(1)
		go producer(i, ch, &producerWg)
	}

	// Start 2 consumers
	for i := 1; i <= 2; i++ {
		consumerWg.Add(1)
		go consumer(i, ch, &consumerWg)
	}

	producerWg.Wait()
	close(ch)
		consumerWg.Wait()
}

func producer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		r := rand.Intn(100)
		fmt.Printf("Producer %d produced: %d\n", id, r)
		ch <- r
	}
}

func consumer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Printf("Consumer %d consumed: %d\n", id, n)
	}
}





11) Producer-Consumer M*N

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sum int    
var sumMutex sync.Mutex
var index int32       

func main() {
	ch := make(chan int)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	m ,n := 3,2 
	for i := 1; i <= m; i++ {
		producerWg.Add(1)
		go producer(i, data, ch, &producerWg)
	}

	for i := 1; i <= n; i++ {
		consumerWg.Add(1)
		go consumer(i, ch, &consumerWg)
	}

	producerWg.Wait()
	close(ch) 
	consumerWg.Wait()
	fmt.Printf("Final sum of consumed values: %d\n", sum)
}








func producer(id int, data []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i := int(atomic.AddInt32(&index, 1)) - 1
		if i >= len(data) {
			break
		}

		value := data[i]
		fmt.Printf("Producer %d produced: %d\n", id, value)
		ch <- value
	}
}

func consumer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	localSum := 0
	for value := range ch {
		fmt.Printf("Consumer %d consumed: %d\n", id, value)
		localSum += value
	}

	sumMutex.Lock()
	sum += localSum
	sumMutex.Unlock()
}






























12)Even Odd 

func main() {
    evenCh := make(chan int)
    oddCh := make(chan int)
    var wg sync.WaitGroup
    
    wg.Add(2)
    go even(evenCh, oddCh, &wg)
    go odd(evenCh, oddCh, &wg)
    
    wg.Wait()
}

func even(evenCh, oddCh chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i:=2; i<=10; i+=2 {
        fmt.Println("Odd :", <-oddCh)
        evenCh<-i
    }
    close(evenCh)
}

func odd(evenCh, oddCh chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i:=1; i<=10; i+=2 {
        oddCh<-i
        fmt.Println("Even :", <-evenCh)
    }
    close(oddCh)
}





13)Global Sum Buffered

package main

import (
	"fmt"
	"sync"
)

var sum int
var mu sync.Mutex // To protect the global sum variable

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	readerCount := 5
	writerCount := 3
	ch := make(chan int, len(data)) // Buffered channel to hold the entire array

	var wgReaders, wgWriters sync.WaitGroup

	// Start reader goroutines
	chunkSize := len(data) / readerCount
	for i := 0; i < readerCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == readerCount-1 { // Handle remaining data in the last chunk
			end = len(data)
		}
		wgReaders.Add(1)
		go reader(ch, data[start:end], &wgReaders)
	}

	// Start writer goroutines
	for i := 0; i < writerCount; i++ {
		wgWriters.Add(1)
		go writer(ch, &wgWriters)
	}

	// Wait for all readers to finish
	wgReaders.Wait()
	close(ch) // Close the channel after all readers are done

	// Wait for all writers to finish processing
	wgWriters.Wait()

	fmt.Println("Total sum:", sum)
}

func reader(ch chan int, data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range data {
		ch <- v // Send each element to the channel
	}
}

func writer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		mu.Lock()
		sum += v // Add each received value to the global sum
		mu.Unlock()
	}
}



14) Global Sum Unbuffered

package main

import (
	"fmt"
	"sync"
)

var sum int
var mu sync.Mutex // To protect the global sum variable

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	readerCount := 5
	writerCount := 3
	ch := make(chan int) // Unbuffered channel

	var wgReaders, wgWriters sync.WaitGroup









	// Start reader goroutines
	chunkSize := len(data) / readerCount
	for i := 0; i < readerCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == readerCount-1 { // Handle remaining data in the last chunk
			end = len(data)
		}
		wgReaders.Add(1)
		go reader(ch, data[start:end], &wgReaders)
	}

	// Start writer goroutines
	for i := 0; i < writerCount; i++ {
		wgWriters.Add(1)
		go writer(ch, &wgWriters)
	}

	// Wait for all readers to finish
	wgReaders.Wait()
	close(ch) // Close the channel after all readers are done

	// Wait for all writers to finish processing
	wgWriters.Wait()

	fmt.Println("Total sum:", sum)
}

func reader(ch chan int, data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range data {
		ch <- v // Send each element to the channel
	}
}

func writer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		mu.Lock()
		sum += v // Add each received value to the global sum
		mu.Unlock()
	}
}