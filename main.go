/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2026-01-13
 * Fileoverview: This program makes a random number 1-10 and prompts the user to guess it until they find it.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	// variables
  var sentence string
  var word string
  var sentenceLower string
  var wordLower string

  reader := bufio.NewReader(os.Stdin)

	// prompts
	fmt.Println("Please enter a sentence? ")
	sentence, _ = reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	fmt.Println("Please enter a word to search for in your sentence? ")
	word, _ = reader.ReadString('\n')
	word = strings.TrimSpace(word)

	// Making all strings lower, to ensure capitalization does not affect
	sentenceLower = strings.ToLower(sentence)
	wordLower = strings.ToLower(word)

	// check
	if strings.Contains(sentenceLower, wordLower) {
		fmt.Printf("Hooray, the word %s was found in the sentence: %s", word, sentence)
	} else {
		fmt.Printf("Sorry, the word %s was not found in the sentence: %s\n", word, sentence)
	}

	fmt.Println("\nDone.")
}
