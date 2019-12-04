// Advent of Code 2019
// Day 2

package main

import(
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Define the opcodes
const (
	ADD = 1
	MULTIPLY = 2
	STOP = 99
)

// A helper function to check for errors

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Update the noun and verb in memory
func updateNounAndVerb(memory []int, noun int, verb int) []int {
	memory[1] = noun
	memory[2] = verb
	return memory
}

// Loads the file into memory and sets the noun and verb

func loadMemory(noun int, verb int) []int {
	// Read the file into a byte slice
	bytes, err := ioutil.ReadFile("day2input.txt")
	check(err)

	// Convert the bytes into a string
	str := string(bytes)

	// Split the string into a string slice on the commas
	strData := strings.Split(str, ",")

	// Convert the string slice into an int slice
	// This is the form we want the memory in
	memory := make([]int, len(strData))
	for i, s := range strData {
		memory[i], err = strconv.Atoi(s)
		check(err)
	}

	// Insert the noun and verb into memory
	updateNounAndVerb(memory, noun, verb)

	// Return the memory
	return memory
}

// Process the instructions stored in memory

func processInstructions(memory []int) int {
	// Place the instruction pointer at the start of memory
	instructionPointer := 0
	stop := false

	// Continue reading instructions until we find a stop
	for !stop {
		// Read the opcode from memory
		opcode := memory[instructionPointer]
		// If we got a STOP opcode, stop processing
		if opcode == STOP {
			stop = true
		} else {
			// Load the input and output addresses from memory
			firstAddress  := memory[instructionPointer + 1]
			secondAddress := memory[instructionPointer + 2]
			outputAddress := memory[instructionPointer + 3]

			// Perform the operation
			switch opcode {
			case ADD:
				memory[outputAddress] = memory[firstAddress] + memory[secondAddress]
			case MULTIPLY:
				memory[outputAddress] = memory[firstAddress] * memory[secondAddress]
			default:
				panic("Unexpected opcode")
			}

			// Increment the instruction pointer
			instructionPointer = instructionPointer + 4

			// Error if the instruction pointer is pointing outside of our memory
			if instructionPointer >= len(memory) {
				panic("Did not found stop opcode")
			}
		}
	}
	
	// Return the output value
	return memory[0]
}

// Find a program code that returns a given output value

func findProgramCode(desiredOutput int) int {
	noun   := 0
	verb   := 0
	output := 0

	// Initialize the memory
	startingMemory := loadMemory(0, 0)

	// Make a copy of memory to avoid having to load it again
	memory := make([]int, len(startingMemory))

	// Loop through the possible noun and verb combinations
	// If we know that the output increases as noun and verb increase, we could speed this up with a binary search
	for noun < 100 {
		for verb < 100 {
			// Restore memory to its original state
			copy(memory, startingMemory)

			// Update the noun and verb in memory
			memory = updateNounAndVerb(memory, noun, verb)
			
			// Process the instructions
			output = processInstructions(memory)

			// If the output matches, finish the loop
			if output == desiredOutput {
				break;
			} else {
				verb++
			}
		}
		// If the output matches, finish this loop too
		if output == desiredOutput {
			break;
		} else {
			noun++
			verb = 0
		}
	}

	// The program code is the noun and verb in a four digit format
	return 100 * noun + verb
}

// Main function

func main() {
	// Puzzle 1
	// Find the output for the given noun and verb
	memory := loadMemory(12, 2)
	output := processInstructions(memory)
	fmt.Println("The output for noun 12 and verb 2 is", output)

	// Puzzle 2
	// Find the noun and verb to get the specified output
	// Program code is defined as 100 * noun + verb
	programCode := findProgramCode(19690720)
	fmt.Println("The program code to get 19690720 is", programCode)
}
