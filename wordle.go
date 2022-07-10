package main

import (
	. "HW3/functions"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var guessWord string
var feedback = "_____"
var feedbackColor = ""
var round = 0

func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(ValidWordList))
	word := strings.ToUpper(ValidWordList[randNum])
	fmt.Println(word, randNum)
	fmt.Println("Welcome to Golang Word-le Project")
	fmt.Println("----------------------------------------")

	// Create map for storing digits and position of digits for guess number and secret number
	// Convert numbers to string for indexing
	wordMap := make(map[string][]int)

	for i := 0; i <= 4; i++ {
		wordMap[word[i:i+1]] = append(wordMap[word[i:i+1]], i)
	}

	for round <= 5 {
		// Input your guess number
		//fmt.Printf("Secret number: %v\n", secretNum)
		//fmt.Println("Guess the Word-le word:\t")
		_, _ = fmt.Scan(&guessWord)
		guessWord := strings.ToUpper(guessWord)

		// Input zero to exit
		if guessWord == "" {
			break
		} else if len([]rune(guessWord)) != 5 {
			err := fmt.Errorf("The word you enter must be five letters! Your input:%+v and input length: %+v", guessWord, len([]rune(guessWord)))
			fmt.Println(err.Error())
			continue
		}

		guessWordMap := make(map[string][]int)

		for i := 0; i <= 4; i++ {
			guessWordMap[guessWord[i:i+1]] = append(guessWordMap[guessWord[i:i+1]], i)
		}

		contains := 0
		position := 0
		for keyGuess := range guessWordMap {
			// If position match occurs
			// Find intersection of position of digit between guess number and secret number
			intersections := Intersection(guessWordMap[keyGuess], wordMap[keyGuess])

			if len(intersections) > 0 {
				position += len(intersections)

				// For correct position digit impose O sign
				for _, index := range intersections {
					feedback = feedback[:index] + "O" + feedback[index+1:]
					feedbackColor = feedbackColor + "<fg=255,255,255;bg=0,170,0;op=underscore;>" + keyGuess + "</>"
				}
				continue
			} else if len(wordMap[keyGuess]) > 0 {
				// If there is no position matches but digit contains in secret number
				// For contains but no position matching digit impose ? sign

				for _, index := range guessWordMap[keyGuess] {
					feedback = feedback[:index] + "?" + feedback[index+1:]
					feedbackColor += "<fg=255,255,255;bg=200,200,0;op=underscore;>" + keyGuess + "</>"
				}

				contains--
			} else {
				feedbackColor += "<fg=255,255,255;op=underscore;>" + keyGuess + "</>"
			}
		}
		//fmt.Print("\033c")
		fmt.Println(feedback)
		fmt.Println(feedbackColor)

		// When find number correctly exit the program
		if position == 5 {
			fmt.Println("----------------------------------------")
			fmt.Printf("Congratulations! You find Wordle word!")
			break
		} else if round == 5 {
			fmt.Println("----------------------------------------")
			fmt.Printf("GAME OVER!")

		}
		// Restart feedback for next round
		feedback = "_____"
		round++
	}
}
