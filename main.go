package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("🎯 Welcome to the Multiple-Choice Quiz Game!")
	fmt.Println("===========================================")
	
	// Display available categories
	categories := GetCategories()
	fmt.Println("\nAvailable Categories:")
	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}
	
	// Get user category choice
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nSelect a category (number): ")
	categoryInput, _ := reader.ReadString('\n')
	categoryInput = strings.TrimSpace(categoryInput)
	categoryIndex, err := strconv.Atoi(categoryInput)
	
	if err != nil || categoryIndex < 1 || categoryIndex > len(categories) {
		fmt.Println("❌ Invalid category selection!")
		return
	}
	
	selectedCategory := categories[categoryIndex-1]
	
	// Start the quiz
	score := StartQuiz(selectedCategory.Questions)
	
	// Display results
	fmt.Printf("\n🎊 Quiz Completed! 🎊\n")
	fmt.Printf("Your Score: %d/%d (%.1f%%)\n", 
		score, len(selectedCategory.Questions), 
		float64(score)/float64(len(selectedCategory.Questions))*100)
}