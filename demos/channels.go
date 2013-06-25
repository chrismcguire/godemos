// Channels

// Borrowed and altered from http://blog.golang.org/advanced-go-concurrency-patterns

// Means of communicating between goroutines
// Blocking by default
// <-channelName to read from a channel
// channelName <- value to write to a channel
package main

import "fmt"
import "time"

func pingpong(value string, channel chan *int) {

    for {
        // Block while waiting to read
        input := <-channel
        fmt.Println(value)
        time.Sleep(100 * time.Millisecond)
        // Block while waiting for another process to read.
        channel <- input
    }

}

func main() {
    // Create a new int channel
    channel := make(chan *int)

    go pingpong("ping", channel)
    go pingpong("pong", channel)
    
    // Write a value to the channel
    gameOn := 1
    channel <- &gameOn

    time.Sleep(1000 * time.Millisecond)
}
