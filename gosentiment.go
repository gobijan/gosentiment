// With gosentiment you pass an english text via cli and it gives you an
// overall rating that indicates if the text positive (>0), neutral (0)
// or negative (<0).
package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

var wordmap map[string]int

func main() {

	// Initialize the wordmap for the rating.
	wordmap = make(map[string]int)

	// Open the ratings file
	file, err := os.Open("AFINN-111-simplified.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split line by spaces
		line := strings.Fields(scanner.Text())
		// fmt.Println(line[0] + " => " + line[1])
		i, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		wordmap[line[0]] = i
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get ready for Command Line
	var text string
	flag.StringVar(&text, "text", "", "specify the text that should be rated.")
	flag.Parse()

	// Calculate the overall rating of the passed text.
	overallRating := RateText(text)
	log.Println("Overall Rating:", overallRating)
}

// Split string by multiple delimitters
// (see: https://groups.google.com/forum/#!topic/golang-nuts/T3ljFVgOwSI)
func delimiters(r rune) bool {
	return r == '.' || r == '!' || r == '?'
}

// SplitTextIntoSentences splits a text into a slice of strings based on punctuation.
func SplitTextIntoSentences(s string) []string {
	sentences := strings.FieldsFunc(s, delimiters)
	return sentences
}

// RateText returns the overall rating for a text that can be composed of several sentences.
func RateText(s string) int {
	sentences := SplitTextIntoSentences(s)

	rating := 0
	for _, sentence := range sentences {
		if len(sentence) > 0 {
			rating += RateSentence(sentence)
		}
	}
	return rating
}

// RateWord gets a word and returns its rating.
func RateWord(s string) int {
	return wordmap[s]
}

// RateSentence receives a sentence and returns its overall rating.
func RateSentence(s string) int {
	// Set intital rating to 0
	rating := 0
	negations := 0
	// Lowercase string
	s = strings.ToLower(s)

	// TestSplitting
	SplitTextIntoSentences(s)

	// Split string into words by blanks
	words := strings.Fields(s)

	for _, element := range words {
		// log.Println(index, element, RateWord(element))
		rating += RateWord(element)
		// println(element)
		if element == "not" {
			negations++
		}
	}

	for i := 0; i < negations; i++ {
		rating *= -1
	}

	log.Println("rating for sentence '", s, "': ", rating)
	return rating
}
