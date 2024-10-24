package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("🎯 Welcome to the Go Quiz Game!")
	fmt.Println("================================")
	
	// Display available categories
	categories := GetCategories()
	for {
		fmt.Println("\nAvailable Categories:")
		for i, category := range categories {
			fmt.Printf("%d. %s\n", i+1, category.Name)
		}
		fmt.Printf("%d. Exit\n", len(categories)+1)
		
		// Get user category choice
		choice := GetUserChoice(len(categories) + 1)
		if choice == len(categories)+1 {
			fmt.Println("Thanks for playing! Goodbye!")
			return
		}
		
		// Start the quiz for selected category
		selectedCategory := categories[choice-1]
		StartQuiz(selectedCategory)
		
		// Ask if user wants to play again
		if !AskToPlayAgain() {
			fmt.Println("Thanks for playing! Goodbye!")
			return
		}
	}
}

func GetUserChoice(max int) int {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Printf("\nChoose a category (1-%d): ", max)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > max {
			fmt.Printf("Invalid choice. Please enter a number between 1 and %d.\n", max)
			continue
		}
		
		return choice
	}
}

func AskToPlayAgain() bool {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("\nWould you like to play again? (y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		
		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		} else {
			fmt.Println("Please enter 'y' for yes or 'n' for no.")
		}
	}
}