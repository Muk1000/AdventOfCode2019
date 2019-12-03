// Advent of Code 2019
// Day 1

package main

import(
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// A helper function to check for errors

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Calculates the fuel needed for the modules

func calculateFuelForModules(includeFuelForFuel bool) int {
    // Open the input file
    file, err := os.Open("day1input.txt")
    check(err)
    defer file.Close()

    // Create a variable to store our total
    totalFuel := 0

    // Scan the file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Convert the line into an integer
        mass, err := strconv.Atoi(scanner.Text())
        check(err)

        // Calculate the fuel
        fuel := calculateFuel(mass, includeFuelForFuel)

        // Add the fuel to our total
        totalFuel += fuel
    }

    // Return the total
    return totalFuel
}

// Calculates the fuel needed for a module

func calculateFuel(mass int, includeFuelForFuel bool) int {
    // Calculate the fuel for the module
    totalFuel := (mass / 3) - 2

    // Check if we need to calculate fuel for the fuel
    currentFuel := totalFuel
    if includeFuelForFuel {
        // Continue until the amount of fuel is negligible
        for currentFuel > 0 {
            // Calculate the fuel for the fuel
            currentFuel = (currentFuel / 3) - 2

            // Add the fuel to our total if it's positive
            if currentFuel > 0 {
                totalFuel += currentFuel
            }
        }
    }

    // Return the total
    return totalFuel
}

// Main function

func main() {
    // Puzzle 1
    // We only care about the fuel for modules
    includeFuelForFuel := false
    fuelForModules := calculateFuelForModules(includeFuelForFuel)
    fmt.Println("The total amount of fuel for the modules is", fuelForModules)

    // Puzzle 2
    // We need the fuel for the modules and the fuel itself
    includeFuelForFuel = true
    totalFuel := calculateFuelForModules(includeFuelForFuel)
    fmt.Println("The total amount of fuel is", totalFuel)
}
