package main

import "fmt"

func main() {

    // Arrays are fixed size
    var cities [3]string
    cities[0] = "Djibouti"
    cities[1] = "Gotham"
    cities[2] = "Philadelphia"
    fmt.Println(cities)
    //cities[3] = "Metropolis"

    numbers := [3]int{4,5,6}
    fmt.Println(numbers)

    // Slices!
    states := []string{"Montana", "Kansas", "Idaho"}
    states = append(states, "Pennsylvania")
    fmt.Println(states)

    // Remove the last element
    states = states[:3]
    fmt.Println(states)

    // Iterating
    for index, state := range states {
        fmt.Println(state + " is at position " + string(index))
    }
}
