package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// I could have just made a map with all alphabet characters
// and their full-width counterparts, but that would be
// the normie way. Instead, I decided to research some
// *very advanced* unicode stuff and figured out that there's
// a constant difference between those characters.
// Honestly, doing it that way probably took me just as much time
// as a map would take, but at least that way I felt smart for a minute.

// Offset between regular latin alphabet
// and relevant full-width characters
const offset = 65248

// This only happens for characters between ! (U+0021 aka 33)
// and ~ (U+007E aka 126), so we'll limit conversion only
// to characters in that range.
const rangeBegin = 33
const rangeEnd = 126

// convertRune applies the offset to a rune, without performing any checks
func convertRune(r rune) rune {
	return r + offset
}

// convertString takes a string and converts regular latin alphabet characters
// to aesthetic full-width characters
func convertString(s string) string {
	convertedString := ""

	// Loop over each character in input string
	for _, char := range s {
		// Perform a check if the character is in the regular offset range
		if char >= rangeBegin && char <= rangeEnd {
			convertedCharacter := convertRune(char)
			convertedString += string(convertedCharacter)
		} else {
			convertedString += string(char)
		}
	}

	return convertedString
}

func main() {
	input := os.Args[1:]

	// Check if argument was provided
	if len(input) > 0 {
		// Parse the string to be converted from args
		parsedInput := strings.Join(input, " ")
		convertedString := convertString(parsedInput)

		// Print the result to console
		fmt.Println(convertedString)
	} else {
		// If not, the program might be running as a pipe
		stat, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			var stdin []string

			// Scan each line
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				stdin = append(stdin, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			
			// Convert and output each line
			for _, line := range stdin {
				convertedLine := convertString(line)
				fmt.Println(convertedLine)
			}
		}
	}

}
