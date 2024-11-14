package main

// Question represents a single quiz question
type Question struct {
	Text    string
	Options []string
	Answer  int // Index of correct answer (0-based)
}

// Category represents a quiz category with related questions
type Category struct {
	Name     string
	Questions []Question
}

// getCategories returns all available quiz categories
func getCategories() []Category {
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
					Answer: 1,
				},
				{
					Text: "Which language is known as the 'mother of all programming languages'?",
					Options: []string{
						"Python",
						"C",
						"Assembly",
						"FORTRAN",
					},
					Answer: 1,
				},
				{
					Text: "What is the time complexity of binary search?",
					Options: []string{
						"O(n)",
						"O(n log n)", 
						"O(log n)",
						"O(1)",
					},
					Answer: 2,
				},
			},
		},
		{
			Name: "Go Language",
			Questions: []Question{
				{
					Text: "Who created the Go programming language?",
					Options: []string{
						"Microsoft",
						"Google",
						"Apple", 
						"Facebook",
					},
					Answer: 1,
				},
				{
					Text: "What is a Goroutine?",
					Options: []string{
						"A type of variable",
						"A lightweight thread", 
						"A package manager",
						"A testing framework",
					},
					Answer: 1,
				},
				{
					Text: "How do you declare a variable in Go?",
					Options: []string{
						"var x int = 5",
						"int x = 5", 
						"x := 5",
						"Both 1 and 3",
					},
					Answer: 3,
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
					Text: "Who painted the Mona Lisa?",
					Options: []string{
						"Vincent van Gogh",
						"Pablo Picasso",
						"Leonardo da Vinci", 
						"Michelangelo",
					},
					Answer: 2,
				},
			},
		},
	}
}