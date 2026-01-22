/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2026-01-13
* @fileoverview This program detects if a word was in the given sentence.
*/

// variables
let sentence: string;
let word: string;
let sentenceLower: string;
let wordLower: string;

// prompt
sentence = prompt("Please enter a sentence?") || "0";
word = prompt("Please enter a word to search for in your sentence?") || "0";

// Making all strings lower, to ensure capitalization does not affect
sentenceLower = sentence.toLowerCase();
wordLower = word.toLowerCase();

// check
if (sentenceLower.includes(wordLower)) {
  console.log(`Hooray, the word ${word} was found in the sentence: ${sentence}`);
} else {
  console.log(`Sorry, the word ${word} was not found in the sentence: ${sentence}`);
}

console.log("\nDone.");
