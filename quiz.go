package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type Category struct {
	Name      string
	Questions []Question
}

func GetCategories() []Category {
	return []Category{
		{
			Name: "Programming Basics",
			Questions: []Question{
				{
					Text: "What does CPU stand for?",
					Options: []string{
						"Computer Processing Unit",
						"Central Processing Unit", 
						"Central Program Utility",
						"Computer Program Unit",
					},
					Answer: 2,
				},
				{
					Text: "Which data structure uses LIFO (Last In First Out) principle?",
					Options: []string{
						"Queue",
						"Stack",
						"Array",
						"Linked List",
					},
					Answer: 2,
				},
				{
					Text: "What is the time complexity of binary search?",
					Options: []string{
						"O(n)",
						"O(n²)",
						"O(log n)",
						"O(1)",
					},
					Answer: 3,
				},
			},
		},
		{
			Name: "Go Language",
			Questions: []Question{
				{
					Text: "How do you declare a variable in Go?",
					Options: []string{
						"var x int",
						"variable x int", 
						"int x",
						"x := int",
					},
					Answer: 1,
				},
				{
					Text: "What is the zero value for a pointer in Go?",
					Options: []string{
						"0",
						"nil",
						"undefined",
						"null",
					},
					Answer: 2,
				},
				{
					Text: "Which keyword is used to create a goroutine?",
					Options: []string{
						"go",
						"routine",
						"goroutine", 
						"async",
					},
					Answer: 1,
				},
			},
		},
		{
			Name: "General Knowledge",
			Questions: []Question{
				{
					Text: "What is the capital of France?",
					Options: []string{
						"London",
						"Berlin", 
						"Paris",
						"Madrid",
					},
					Answer: 3,
				},
				{
					Text: "Which planet is known as the Red Planet?",
					Options: []string{
						"Venus",
						"Mars",
						"Jupiter",
						"Saturn",
					},
					Answer: 2,
				},
				{
					Text: "How many continents are there?",
					Options: []string{
						"5",
						"6",
						"7",
						"8",
					},
					Answer: 3,
				},
			},
		},
	}
}

func StartQuiz(category Category) {
	fmt.Printf("\n🏁 Starting %s Quiz!\n", category.Name)
	fmt.Println(strings.Repeat("=", 40))
	
	score := 0
	totalQuestions := len(category.Questions)
	reader := bufio.NewReader(os.Stdin)
	
	for i, question := range category.Questions {
		fmt.Printf("\nQuestion %d/%d:\n", i+1, totalQuestions)
		fmt.Printf("❓ %s\n", question.Text)
		
		// Display options
		for j, option := range question.Options {
			fmt.Printf("   %d. %s\n", j+1, option)
		}
		
		// Get user answer
		userAnswer := GetAnswer(len(question.Options), reader)
		
		// Check if answer is correct
		if userAnswer == question.Answer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Printf("❌ Incorrect! The correct answer was: %d. %s\n", 
				question.Answer, question.Options[question.Answer-1])
		}
	}
	
	// Display final score
	DisplayScore(score, totalQuestions)
}

func GetAnswer(maxOptions int, reader *bufio.Reader) int {
	for {
		fmt.Printf("\nEnter your answer (1-%d): ", maxOptions)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		answer, err := strconv.Atoi(input)
		if err != nil || answer < 1 || answer > maxOptions {
			fmt.Printf("Invalid answer. Please enter a number between 1 and %d.\n", maxOptions)
			continue
		}
		
		return answer
	}
}

func DisplayScore(score, total int) {
	percentage := float64(score) / float64(total) * 100
	
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Printf("🎯 Quiz Completed!\n")
	fmt.Printf("📊 Your Score: %d/%d (%.1f%%)\n", score, total, percentage)
	
	switch {
	case percentage >= 90:
		fmt.Println("🏆 Excellent! You're a quiz master!")
	case percentage >= 70:
		fmt.Println("👍 Great job! You know your stuff!")
	case percentage >= 50:
		fmt.Println("😊 Good effort! Keep learning!")
	default:
		fmt.Println("💪 Don't give up! Try again!")
	}
}