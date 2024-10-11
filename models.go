package main

// Question represents a single multiple-choice question
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