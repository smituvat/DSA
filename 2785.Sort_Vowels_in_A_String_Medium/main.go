package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to check if a character is a vowel
func isVowel(char rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, char)
}

// Function to sort the vowels in a string in ascending order
func sortVowels(input string) string {
	var vowels []rune

	// Extract vowels from the input string
	for _, char := range input {
		if isVowel(char) {
			vowels = append(vowels, char)
		}
	}

	// Sort the vowels in ascending order
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	// Replace the vowels in the input string with the sorted vowels
	vowelIndex := 0
	output := strings.Map(func(char rune) rune {
		if isVowel(char) {
			vowel := vowels[vowelIndex]
			vowelIndex++
			return vowel
		}
		return char
	}, input)

	return output
}

func main() {
	inputString := "hello world"
	sortedString := sortVowels(inputString)
	fmt.Println("Input String:", inputString)
	fmt.Println("Sorted Vowels:", sortedString)
}
