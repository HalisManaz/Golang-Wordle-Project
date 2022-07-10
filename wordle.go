package main

import (
	"fmt"
	"github.com/gookit/color"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var guessWord string
var feedbackColor = ""
var round = 0

func main() {
	// Create number random and select random words in list
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(ValidWordList))
	word := strings.ToUpper(ValidWordList[randNum])
	fmt.Println(word, randNum)
	fmt.Println("Welcome to Golang Word-le Project")
	fmt.Println("----------------------------------------")

	// Create map for storing letter and position of letters for Word-le word
	wordMap := make(map[string][]int)

	for i := 0; i <= 4; i++ {
		wordMap[word[i:i+1]] = append(wordMap[word[i:i+1]], i)
	}

	for round < 5 {
		// Input your guess number
		//fmt.Println("Guess the Word-le word:\t")
		_, _ = fmt.Scanln(&guessWord)
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		err := c.Run()
		if err != nil {
			return
		}
		//color.Println("<fg=255,255,255;>" + "_____" + "</>")

		guessWord := strings.ToUpper(guessWord)

		// Input XXXXX to exit
		if guessWord == "XXXXX" {
			break
		} else if len([]rune(guessWord)) != 5 {
			// Check input length is equal to 5 or not
			err := fmt.Errorf("the word you enter must be five letters Your input:%+v and input length: %+v", guessWord, len([]rune(guessWord)))
			fmt.Println(err.Error())
			continue
		}

		// Create map for storing letter and position of letters for guess word
		guessWordMap := make(map[string][]int)

		for i := 0; i <= 4; i++ {
			guessWordMap[guessWord[i:i+1]] = append(guessWordMap[guessWord[i:i+1]], i)
		}

		// Create default feedback with non-color guess word
		feedbackColorMaps := make(map[int]string)
		for keyGuess, indexes := range guessWordMap {
			for i := 0; i < len(indexes); i++ {
				add := "<fg=255,255,255;op=underscore;>" + keyGuess + "</>"
				feedbackColorMaps[i] = add
			}
		}

		position := 0
		for keyGuess := range guessWordMap {

			comparisonMap := map[int]int{}

			// Find intersection and difference of index of letters of guess word
			intersections := Intersection(guessWordMap[keyGuess], wordMap[keyGuess])
			diff := Difference(guessWordMap[keyGuess], wordMap[keyGuess])
			diff2 := Difference(wordMap[keyGuess], guessWordMap[keyGuess])

			// Add intersection values
			for i := 0; i < len(intersections); i++ {
				comparisonMap[intersections[i]] = intersections[i]
			}

			// Add difference values
			for i := 0; i < len(diff); i++ {
				if len(diff2) == 0 || i >= len(diff2) {
					comparisonMap[diff[i]] = -1
				} else {
					comparisonMap[diff[i]] = diff2[i]
				}
			}

			for guessIndex, wordIndex := range comparisonMap {
				if wordIndex == -1 {
					// If not contains
					add := "<fg=255,255,255;op=underscore,bold;>" + keyGuess + "</>"
					feedbackColorMaps[guessIndex] = add
					continue
				} else if wordIndex != guessIndex {
					// Contains but position wrong
					add := "<fg=255,255,255;bg=200,200,0;op=underscore,bold;>" + keyGuess + "</>"
					feedbackColorMaps[guessIndex] = add
				} else if wordIndex == guessIndex {
					// Contains and position are correct
					add := "<fg=255,255,255;bg=0,170,0;op=underscore,bold;>" + keyGuess + "</>"
					feedbackColorMaps[guessIndex] = add
					position++
				}
			}
		}

		for i := 0; i <= 4; i++ {
			// Create colorful feedback
			feedbackColor += feedbackColorMaps[i]
		}

		color.Println(feedbackColor)
		feedbackColor += "\n"

		// When find Word-le words correctly exit the program
		if position == 5 {
			fmt.Println("----------------------------------------")
			fmt.Printf("Congratulations! You find Wordle word!")
			break
		} else if round == 4 {
			fmt.Println("----------------------------------------")
			fmt.Printf("GAME OVER!")

		}
		// Restart feedback for next round
		//feedbackColor = ""
		round++
	}
}

func Intersection(first, second []int) []int {
	var intersections []int

	for _, i := range first {
		for _, j := range second {
			if i == j {
				intersections = append(intersections, i)
			}
		}
	}
	return intersections
}

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []int) []int {
	mb := make(map[int]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
