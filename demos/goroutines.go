// Goroutines

// Independently executing function with its own call stack (technically Communicating Sequential Processes)
// Technically not a thread, but basically a cheap thread...
// Super lightweight. Practical to have thousand or hundreds of thousands of goroutines.
package main

import "fmt"
import "time"

func sleepAndPrint(sleepTime int, output string) {
	// Wait sleepTime ms, then print output
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println(output)
}

func main() {
	// Spawn 2 goroutines
	go sleepAndPrint(500, "let the dogs out?")
	go sleepAndPrint(100, "Who")

	// If the main function finishes, it closes the gorouites and exits
	time.Sleep(3000 * time.Millisecond)
}
