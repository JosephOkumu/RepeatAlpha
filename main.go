package main

import (
    "fmt"
    "os"
)

func main() {
    // Check if the number of arguments is correct
    if len(os.Args) != 2 {
        return
    }

    // Get the input string from the command-line argument
    input := os.Args[1]
    result := ""

    // Iterate over each character in the input string
    for _, char := range input {
        // Check if the character is a letter
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            // Calculate the index of the character in the alphabet (1-based)
            index := int(char) - (int('a') - 1)
            if char >= 'A' && char <= 'Z' {
                index = int(char) - (int('A') - 1)
            }

            // Repeat the character the specified number of times
            for i := 0; i < index; i++ {
                result += string(char)
            }
        } else {
            // If the character is not a letter, add it to the result as is
            result += string(char)
        }
    }

    // Add a newline character to the result
    result += "\n"

    // Write the result to the standard output
    fmt.Print(result)
}
