// Arrays, Slices, and Maps

// The hardest change coming from python

package main

import "fmt"
//import "strconv"

func main() {

    // Arrays are fixed size
    var cities [3]string
    cities[0] = "Djibouti"
    cities[1] = "Gotham"
    cities[2] = "Philadelphia"
    fmt.Println(cities)
    // This kills the program...
    //cities[3] = "Metropolis"

//    numbers := [3]int{4,5,6}
//    fmt.Println(numbers)
//
//    // Slices! Just use them instead.
//    states := []string{"Montana", "Kansas", "Idaho"}
//    states = append(states, "Pennsylvania")
//    fmt.Println(states)
//
//    _ = make([]int, 5)
//
//    // Remove the last element
//    states = states[:3]
//    fmt.Println(states)
//
//    // Iterating
//    for index, state := range states {
//        // fix index printing
//        fmt.Println(state + " is at position " + strconv.Itoa(index))
//    }
//
//    // Maps (dictionaries)
//    legs := make(map[string]int)
//    legs["human"] = 2
//    legs["insect"] = 6
//    legs["chair"] = 4
//    fmt.Println(legs)
//    delete(legs, "chair")
//    fmt.Println(legs)

}
