// Advent of Code 2019
// Day 4

package main

import(
    "fmt"
    "math"
)

// A helper function to check for errors
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Find the number of possible passwords in the given range
func findNumberOfPossiblePasswords(min int, max int, allowGroups bool) int {
    // Check each number
    numberOfPossiblePasswords := 0
    for num := min; num <= max; num++ {
        foundRepeat := false
        repeatCounter := 1
        previousDigit := -1
        // Check each digit of the number
        for i := 5; i >= 0; i-- {
            currentDigit := (num / int(math.Pow10(i))) % 10
            // If we have a previous digit, make the checks
            if previousDigit >= 0 {
                if previousDigit > currentDigit {
                    // The previous digit can't be larger
                    // We could improve this by skipping ahead until this digit is no longer smaller
                    foundRepeat = false
                    break
                } else if previousDigit == currentDigit {
                    if allowGroups {
                        // If groups are allowed, this is sufficient
                        foundRepeat = true
                    } else if foundRepeat == false {
                        // Increment the repeat counter
                        repeatCounter++
                        // If this is the last digit, check if it ended with a group of two
                        if i == 0 {
                            if repeatCounter == 2 {
                                foundRepeat = true
                            }
                        }
                    }
                } else {
                    // The numbers are different, so check if we found a group of two
                    if foundRepeat == false {
                        if repeatCounter == 2 {
                            foundRepeat = true
                        }
                        // Reset the counter
                        repeatCounter = 1
                    }
                }
            }
            // Make the current digit the new previous digit
            previousDigit = currentDigit
        }

        // If the number had a valid repeat, increment the counter
        if foundRepeat {
            numberOfPossiblePasswords++
        }
    }

    return numberOfPossiblePasswords
}

// Main function
func main() {
    // Define the range of possible values
    min := 235741
    max := 706948

    // Puzzle 1
    num := findNumberOfPossiblePasswords(min, max, true)
    fmt.Println("The number of possible passwords when allowing large groups of matching digits is", num)

    // Puzzle 2
    num = findNumberOfPossiblePasswords(min, max, false)
    fmt.Println("The number of possible passwords when not allowing large groups of matching digits is", num)
}
