package main

import (
	"fmt"
	// "sync"
	"time"
)
package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup

// func worker(id int) {
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d done\n", id)
// }

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			worker(id)
// 		}(i) // Pass `i` as an argument to the goroutine
// 	}

// 	wg.Wait() // Wait for all goroutines to finish
// }

// var wg sync.WaitGroup

// func worker(id int, wg sync.WaitGroup) {
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d done\n", id)

// }

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go func() {
//]
// 			worker(i, wg)
// 		}() // Pass `i` as an argument to the goroutine
// 	}

// 	 wg.Wait() // Wait for all goroutines to finish
// }





func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

}

func main() {
	for i := 1; i <= 5; i++ {
	
		go func() {
		
			worker(i)
		}()
	
	}

	time.Sleep(5 * time.Second)
}
