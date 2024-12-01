package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Question represents a single quiz question
type Question struct {
	Text    string
	Options []string
	Answer  int // Index of correct answer (0-based)
}

// Quiz represents a collection of questions for a specific topic
type Quiz struct {
	Name      string
	Questions []Question
}

// LoadQuizzes loads all available quiz topics
func LoadQuizzes() []Quiz {
	return []Quiz{
		{
			Name: "Go Programming Basics",
			Questions: []Question{
				{
					Text: "What keyword is used to declare a variable in Go?",
					Options: []string{
						"var",
						"let",
						"const",
						"variable",
					},
					Answer: 0,
				},
				{
					Text: "Which of these is NOT a primitive type in Go?",
					Options: []string{
						"int",
						"string",
						"float",
						"class",
					},
					Answer: 3,
				},
				{
					Text: "How do you create a slice in Go?",
					Options: []string{
						"make([]int, 5)",
						"new([]int, 5)",
						"slice([]int, 5)",
						"create([]int, 5)",
					},
					Answer: 0,
				},
			},
		},
		{
			Name: "Computer Science Fundamentals",
			Questions: []Question{
				{
					Text: "What does CPU stand for?",
					Options: []string{
						"Central Processing Unit",
						"Computer Processing Unit",
						"Central Program Unit",
						"Computer Program Unit",
					},
					Answer: 0,
				},
				{
					Text: "Which data structure uses LIFO (Last In First Out)?",
					Options: []string{
						"Queue",
						"Stack",
						"Array",
						"Linked List",
					},
					Answer: 1,
				},
				{
					Text: "What is the time complexity of binary search?",
					Options: []string{
						"O(1)",
						"O(n)",
						"O(log n)",
						"O(n²)",
					},
					Answer: 2,
				},
			},
		},
		{
			Name: "General Knowledge",
			Questions: []Question{
				{
					Text: "What is the capital of Japan?",
					Options: []string{
						"Seoul",
						"Beijing",
						"Tokyo",
						"Bangkok",
					},
					Answer: 2,
				},
				{
					Text: "Which planet is known as the Red Planet?",
					Options: []string{
						"Venus",
						"Mars",
						"Jupiter",
						"Saturn",
					},
					Answer: 1,
				},
				{
					Text: "Who wrote 'Romeo and Juliet'?",
					Options: []string{
						"Charles Dickens",
						"William Shakespeare",
						"Jane Austen",
						"Mark Twain",
					},
					Answer: 1,
				},
			},
		},
	}
}

// StartQuiz runs the quiz and returns the score
func StartQuiz(quiz Quiz) int {
	fmt.Printf("\n📝 Starting Quiz: %s\n", quiz.Name)
	fmt.Println("========================")
	
	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	
	for i, question := range quiz.Questions {
		fmt.Printf("\nQuestion %d/%d: %s\n", i+1, len(quiz.Questions), question.Text)
		
		// Display options
		for j, option := range question.Options {
			fmt.Printf("  %d. %s\n", j+1, option)
		}
		
		// Get user answer
		var answer int
		for {
			fmt.Print("Your answer (number): ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			
			_, err := fmt.Sscanf(input, "%d", &answer)
			if err != nil || answer < 1 || answer > len(question.Options) {
				fmt.Printf("❌ Please enter a number between 1 and %d\n", len(question.Options))
				continue
			}
			break
		}
		
		// Check if answer is correct
		if answer-1 == question.Answer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Printf("❌ Incorrect! The correct answer was: %d. %s\n", 
				question.Answer+1, question.Options[question.Answer])
		}
	}
	
	return score
}