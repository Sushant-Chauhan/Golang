package main

import (
	"fmt"
	"sync"
)

// func firstfunc() {
// 	fmt.Println("First function")
// }

// func secondfunc() {
// 	fmt.Println("Second function")
// }

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	//some task
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Println("Main function started..")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Main function ended..")

	// go firstfunc()
	// go secondfunc()
}
