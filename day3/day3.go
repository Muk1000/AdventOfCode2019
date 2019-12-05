// Advent of Code 2019
// Day 3

package main

import(
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

// Location stores a position on the grid
type Location struct {
    x int
    y int
}

// WireSegment stores the location of a wire segment and how many steps it took to get there
type WireSegment struct {
    location Location
    steps int
}

// Define the directions
const (
    RIGHT = 82
    LEFT  = 76
    UP    = 85
    DOWN  = 68
)

// A helper function to check for errors
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// A helper function to get the absolute value of ints
func abs(num int) int {
    if num < 0 {
        num = num * -1
    }
    return num
}

// Check for a wire intersection
func checkForIntersection(firstWirePath map[Location]WireSegment, secondWirePath map[Location]WireSegment, position Location, minDistance int, minWireSteps int) (int, int) {
    // Check if the first wire also crossed this position
    if firstWireSegment, ok := firstWirePath[position]; ok {
        // Update the distance if it's smaller
        intersectionDistance := abs(position.x) + abs(position.y)
        if intersectionDistance < minDistance {
            minDistance = intersectionDistance
        }
        // Update the number of steps if it's smaller
        wireSteps := firstWireSegment.steps + secondWirePath[position].steps
        if wireSteps < minWireSteps {
            minWireSteps = wireSteps
        }
    }

    return minDistance, minWireSteps
}

// Find the closest place the wires intersect
func findIntersections() (int, int) {
    // Open the input file
    file, err := os.Open("day3input.txt")
    check(err)
    defer file.Close()

    ////////////
    // Wire 1 //
    ////////////

    // Read the first wire's instructions into a slice of strings
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    wireDirections := strings.Split(scanner.Text(), ",")

    // Set our starting position and step count
    position := Location{0, 0}
    steps    := 0

    // Create a hash to store the path of the first wire
    firstWirePath := make(map[Location]WireSegment)
    firstWirePath[position] = WireSegment{position, steps}

    // Process the instructions for the first wire
    for i := 0; i < len(wireDirections); i++ {
        // Split the instruction into direction and distance
        vector    := wireDirections[i]
        direction := vector[0]
        distance, err := strconv.Atoi(vector[1:len(vector)])
        check(err)
        
        // Move along the wire
        switch direction {
        case RIGHT:
            for j := 0; j < distance; j++ {
                // Increment the position in the correct direction
                position.x++
                // Increment the step counter
                steps++
                // If we haven't recorded this position yet, add it to our path
                if _, ok := firstWirePath[position]; ok {} else {
                    firstWirePath[position] = WireSegment{position, steps}
                }
            }
        case LEFT:
            for j := 0; j < distance; j++ {
                position.x--
                steps++
                if _, ok := firstWirePath[position]; ok {} else {
                    firstWirePath[position] = WireSegment{position, steps}
                }
            }
        case UP:
            for j := 0; j < distance; j++ {
                position.y++
                steps++
                if _, ok := firstWirePath[position]; ok {} else {
                    firstWirePath[position] = WireSegment{position, steps}
                }
            }
        case DOWN:
            for j := 0; j < distance; j++ {
                position.y--
                steps++
                if _, ok := firstWirePath[position]; ok {} else {
                    firstWirePath[position] = WireSegment{position, steps}
                }
            }
        default:
            panic("Invalid direction")
        }
    }

    ////////////
    // Wire 2 //
    ////////////

    // Read the second wire's instructions into a slice of strings
    scanner.Scan()
    wireDirections = strings.Split(scanner.Text(), ",")

    // Start our minimum distance and steps at the max int
    minDistance  := math.MaxInt32
    minWireSteps := math.MaxInt32

    // Reset our position and step counter
    position = Location{0, 0}
    steps = 0

    // Create a hash to store the path of the second wire
    secondWirePath := make(map[Location]WireSegment)
    secondWirePath[position] = WireSegment{position, steps}

    // Process the instructions for the second wire
    for i := 0; i < len(wireDirections); i++ {
        // Split the instruction into direction and distance
        vector    := wireDirections[i]
        direction := vector[0]
        distance, err := strconv.Atoi(vector[1:len(vector)])
        check(err)
        
        // Move along the wire
        switch direction {
        case RIGHT:
            for j := 0; j < distance; j++ {
                // Increment the position in the correct direction
                position.x++
                // Increment the step counter
                steps++
                // If we haven't recorded this position yet, add it to our path
                if _, ok := secondWirePath[position]; ok {} else {
                    secondWirePath[position] = WireSegment{position, steps}
                }
                // Check if we've found an intersection and update our minimum values
                minDistance, minWireSteps = checkForIntersection(firstWirePath, secondWirePath, position, minDistance, minWireSteps)
            }
        case LEFT:
            for j := 0; j < distance; j++ {
                position.x--
                steps++
                if _, ok := secondWirePath[position]; ok {} else {
                    secondWirePath[position] = WireSegment{position, steps}
                }
                minDistance, minWireSteps = checkForIntersection(firstWirePath, secondWirePath, position, minDistance, minWireSteps)
            }
        case UP:
            for j := 0; j < distance; j++ {
                position.y++
                steps++
                if _, ok := secondWirePath[position]; ok {} else {
                    secondWirePath[position] = WireSegment{position, steps}
                }
                minDistance, minWireSteps = checkForIntersection(firstWirePath, secondWirePath, position, minDistance, minWireSteps)
            }
        case DOWN:
            for j := 0; j < distance; j++ {
                position.y--
                steps++
                if _, ok := secondWirePath[position]; ok {} else {
                    secondWirePath[position] = WireSegment{position, steps}
                }
                minDistance, minWireSteps = checkForIntersection(firstWirePath, secondWirePath, position, minDistance, minWireSteps)
            }
        default:
            panic("Invalid direction")
        }
    }

    return minDistance, minWireSteps
}

// Main function
func main() {
    distance, wireSteps := findIntersections()
    fmt.Println("The nearest intersection of wires is", distance, "steps away")
    fmt.Println("The minimum steps along the wires to reach an intersection is", wireSteps)
}
