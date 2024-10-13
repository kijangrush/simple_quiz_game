package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QuizManager struct {
	categories map[string][]Question
}

type Question struct {
	Text    string
	Options []string
	Answer  int
}

func NewQuizManager() *QuizManager {
	qm := &QuizManager{
		categories: make(map[string][]Question),
	}
	qm.initializeQuestions()
	return qm
}

func (qm *QuizManager) initializeQuestions() {
	// Programming Questions
	qm.categories["Programming"] = []Question{
		{
			Text: "What does 'GOROOT' environment variable point to in Go?",
			Options: []string{
				"Go source code location",
				"Go binary installation directory",
				"Go workspace directory",
				"Go module cache",
			},
			Answer: 1,
		},
		{
			Text: "Which keyword is used to create a new goroutine in Go?",
			Options: []string{
				"go",
				"goroutine",
				"async",
				"spawn",
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
	}

	// Science Questions
	qm.categories["Science"] = []Question{
		{
			Text: "What is the chemical symbol for gold?",
			Options: []string{
				"Go",
				"Gd",
				"Au",
				"Ag",
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
	}

	// History Questions
	qm.categories["History"] = []Question{
		{
			Text: "In which year did World War II end?",
			Options: []string{
				"1944",
				"1945",
				"1946",
				"1947",
			},
			Answer: 1,
		},
		{
			Text: "Who was the first President of the United States?",
			Options: []string{
				"Thomas Jefferson",
				"George Washington",
				"Abraham Lincoln",
				"John Adams",
			},
			Answer: 1,
		},
	}
}

func (qm *QuizManager) DisplayCategories() {
	fmt.Println("\n📚 Available Categories:")
	fmt.Println("----------------------")
	
	i := 1
	for category, questions := range qm.categories {
		fmt.Printf("%d. %s (%d questions)\n", i, category, len(questions))
		i++
	}
}

func (qm *QuizManager) StartQuiz(category string) {
	questions, exists := qm.categories[category]
	if !exists {
		fmt.Printf("❌ Category '%s' not found!\n", category)
		return
	}

	fmt.Printf("\n🚀 Starting %s Quiz!\n", category)
	fmt.Println("======================")

	score := 0
	totalQuestions := len(questions)
	reader := bufio.NewReader(os.Stdin)

	for i, question := range questions {
		fmt.Printf("\nQuestion %d/%d:\n", i+1, totalQuestions)
		fmt.Printf("❓ %s\n\n", question.Text)
		
		// Display options
		for j, option := range question.Options {
			fmt.Printf("  %d. %s\n", j+1, option)
		}
		
		// Get user answer
		userAnswer := getUserAnswer(len(question.Options), reader)
		
		// Check if answer is correct (convert to 0-based index)
		if userAnswer-1 == question.Answer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Printf("❌ Incorrect! The correct answer was: %d. %s\n", 
				question.Answer+1, question.Options[question.Answer])
		}
	}

	// Display final score
	qm.displayScore(score, totalQuestions)
}

func getUserAnswer(maxOptions int, reader *bufio.Reader) int {
	for {
		fmt.Printf("\nEnter your answer (1-%d): ", maxOptions)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		answer, err := strconv.Atoi(input)
		if err != nil || answer < 1 || answer > maxOptions {
			fmt.Printf("❌ Please enter a number between 1 and %d\n", maxOptions)
			continue
		}
		return answer
	}
}

func (qm *QuizManager) displayScore(score, total int) {
	fmt.Println("\n🎉 Quiz Completed!")
	fmt.Println("=================")
	fmt.Printf("Final Score: %d/%d\n", score, total)
	
	percentage := float64(score) / float64(total) * 100
	fmt.Printf("Percentage: %.1f%%\n", percentage)
	
	switch {
	case percentage >= 90:
		fmt.Println("🏆 Excellent! You're a quiz master!")
	case percentage >= 70:
		fmt.Println("👍 Great job! You know your stuff!")
	case percentage >= 50:
		fmt.Println("😊 Good effort! Keep learning!")
	default:
		fmt.Println("📚 Keep studying! You'll do better next time!")
	}
}