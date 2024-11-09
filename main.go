package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🎯 Welcome to the Multiple-Choice Quiz Game!")
	fmt.Println("===========================================")
	
	// Display available categories
	fmt.Println("\nAvailable Categories:")
	for i, category := range getCategories() {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}
	
	// Get user category choice
	var choice int
	fmt.Print("\nSelect a category (1-3): ")
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > 3 {
		fmt.Println("❌ Invalid choice. Please run the program again.")
		return
	}
	
	// Run the quiz
	categories := getCategories()
	selectedCategory := categories[choice-1]
	
	fmt.Printf("\nStarting %s Quiz...\n", selectedCategory.Name)
	fmt.Println("========================")
	
	score := runQuiz(selectedCategory.Questions)
	
	// Display results
	fmt.Printf("\n🎊 Quiz Completed! 🎊\n")
	fmt.Printf("Your Score: %d/%d (%.1f%%)\n", 
		score, len(selectedCategory.Questions), 
		float64(score)/float64(len(selectedCategory.Questions))*100)
	
	if score == len(selectedCategory.Questions) {
		fmt.Println("🏆 Perfect score! Excellent work!")
	} else if score >= len(selectedCategory.Questions)/2 {
		fmt.Println("👍 Good job! Keep practicing!")
	} else {
		fmt.Println("💪 Keep learning! You'll do better next time!")
	}
}