package main

import (
	"fmt"
)

// runQuiz executes the quiz and returns the score
func runQuiz(questions []Question) int {
	score := 0
	
	for i, question := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Text)
		
		// Display options
		for j, option := range question.Options {
			fmt.Printf("  %c. %s\n", 'A'+j, option)
		}
		
		// Get user answer
		var userAnswer string
		fmt.Print("Your answer (A/B/C/D): ")
		fmt.Scan(&userAnswer)
		
		// Convert answer to index (A=0, B=1, C=2, D=3)
		answerIndex := convertAnswerToIndex(userAnswer)
		
		if answerIndex == question.Answer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Printf("❌ Incorrect! The correct answer was %c. %s\n", 
				'A'+question.Answer, question.Options[question.Answer])
		}
	}
	
	return score
}

// convertAnswerToIndex converts letter answer to index (A=0, B=1, etc.)
func convertAnswerToIndex(answer string) int {
	if len(answer) == 0 {
		return -1
	}
	
	switch answer[0] {
	case 'A', 'a':
		return 0
	case 'B', 'b':
		return 1
	case 'C', 'c':
		return 2
	case 'D', 'd':
		return 3
	default:
		return -1
	}
}