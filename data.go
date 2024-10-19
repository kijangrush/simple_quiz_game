package main

// GetCategories returns all available quiz categories
func GetCategories() []Category {
	return []Category{
		{
			Name: "Programming Fundamentals",
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
					Text: "Which data structure uses LIFO (Last In First Out) principle?",
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
			Name: "Go Language",
			Questions: []Question{
				{
					Text: "Which keyword is used to declare a variable in Go?",
					Options: []string{
						"var",
						"let",
						"const",
						"variable",
					},
					Answer: 0,
				},
				{
					Text: "What is the zero value of a pointer in Go?",
					Options: []string{
						"0",
						"nil", 
						"undefined",
						"null",
					},
					Answer: 1,
				},
				{
					Text: "How do you create a new goroutine?",
					Options: []string{
						"go functionName()",
						"async functionName()",
						"thread functionName()",
						"routine functionName()",
					},
					Answer: 0,
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
			},
		},
	}
}