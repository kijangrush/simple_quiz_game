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
	
	// Initialize quiz manager
	quizManager := NewQuizManager()
	
	// Display available categories
	quizManager.DisplayCategories()
	
	// Get user category choice
	category := getUserCategoryChoice(quizManager)
	
	// Start the quiz
	quizManager.StartQuiz(category)
}

func getUserCategoryChoice(qm *QuizManager) string {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("\nChoose a category (enter number): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(qm.categories) {
			fmt.Println("❌ Invalid choice. Please try again.")
			continue
		}
		
		// Convert choice to category name
		i := 1
		for cat := range qm.categories {
			if i == choice {
				return cat
			}
			i++
		}
	}
}