package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🎯 Welcome to the Go Quiz Game!")
	fmt.Println("===============================")
	
	// Load quiz data
	quizzes := LoadQuizzes()
	
	// Display available topics
	fmt.Println("\nAvailable Quiz Topics:")
	for i, topic := range quizzes {
		fmt.Printf("%d. %s\n", i+1, topic.Name)
	}
	
	// Get user topic choice
	var choice int
	fmt.Print("\nSelect a topic (number): ")
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(quizzes) {
		fmt.Println("❌ Invalid choice!")
		return
	}
	
	selectedQuiz := quizzes[choice-1]
	
	// Start the quiz
	score := StartQuiz(selectedQuiz)
	
	// Display results
	fmt.Printf("\n🎉 Quiz Completed!\n")
	fmt.Printf("Your score: %d/%d (%.1f%%)\n", 
		score, len(selectedQuiz.Questions), 
		float64(score)/float64(len(selectedQuiz.Questions))*100)
	
	if float64(score)/float64(len(selectedQuiz.Questions)) >= 0.7 {
		fmt.Println("🏆 Excellent! You passed!")
	} else {
		fmt.Println("💪 Keep practicing!")
	}
}