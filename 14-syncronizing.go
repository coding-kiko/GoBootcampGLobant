package main

import (
	"fmt"
	"sync"
	"time"
)

func racer(tag int, wg *sync.WaitGroup) {
	defer wg.Done() // after the routine ends, I can tell the WaitGroup to decrement the wait count by 1
	fmt.Printf("Racer %d starts sprinting\n", tag)
	time.Sleep(2 * time.Second) // at this point the 'racers' will be running concurrently
	fmt.Printf("Racer %d reaches the finish line\n", tag)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ { // 5 racers are created, which will be all 'running' on different goroutines
		wg.Add(1) // for every new goroutine I add +1
		go racer(i, &wg)
	}

	wg.Wait() // this waits for all routines asigned to the same waitgroup to finish, else the main routine
		  // would finish and spoil the party
}

// output:       NOTICE: the result of the 'race' will be different each time
// Racer 1 starts sprinting
// Racer 5 starts sprinting
// Racer 4 starts sprinting
// Racer 2 starts sprinting
// Racer 3 starts sprinting
// Racer 3 reaches the finish line
// Racer 2 reaches the finish line
// Racer 4 reaches the finish line
// Racer 1 reaches the finish line
// Racer 5 reaches the finish line
