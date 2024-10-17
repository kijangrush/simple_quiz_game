package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StartQuiz runs the quiz and returns the score
func StartQuiz(questions []Question) int {
	score := 0
	reader := bufio.NewReader(os.Stdin)
	
	for i, question := range questions {
		fmt.Printf("\nQuestion %d/%d:\n", i+1, len(questions))
		fmt.Printf("❓ %s\n", question.Text)
		
		// Display options
		for j, option := range question.Options {
			fmt.Printf("  %d. %s\n", j+1, option)
		}
		
		// Get user answer
		fmt.Print("Your answer (number): ")
		answerInput, _ := reader.ReadString('\n')
		answerInput = strings.TrimSpace(answerInput)
		userAnswer, err := strconv.Atoi(answerInput)
		
		if err != nil {
			fmt.Println("❌ Invalid input! Please enter a number.")
			continue
		}
		
		// Check if answer is correct (convert to 0-based index)
		if userAnswer-1 == question.Answer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Printf("❌ Wrong! The correct answer was: %d. %s\n", 
				question.Answer+1, question.Options[question.Answer])
		}
	}
	
	return score
}