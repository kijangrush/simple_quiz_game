# Multiple-Choice Quiz Game

A simple command-line multiple-choice quiz game written in Go with scoring and topic categories.

## Features

- 🎯 Multiple topic categories (Programming, Science, History)
- 📊 Real-time scoring and progress tracking
- 🏆 Performance feedback based on score
- ❌ Error handling for invalid inputs
- 📝 Easy to extend with new questions and categories

## How to Use

### Prerequisites
- Go 1.21 or later installed on your system

### Installation & Running

1. **Create a new directory for the project:**
   ```bash
   mkdir quiz-game
   cd quiz-game
   ```

2. **Create the files:**
   - Copy the code from above into `main.go`, `quiz_manager.go`, and `go.mod`

3. **Run the game:**
   ```bash
   go run main.go quiz_manager.go
   ```

### Game Flow

1. **Start the game:** The program will display a welcome message and available categories
2. **Choose a category:** Select a category by entering the corresponding number
3. **Answer questions:** Read each question and select your answer by entering 1, 2, 3, or 4
4. **Get immediate feedback:** After each question, you'll see if you answered correctly
5. **View final score:** At the end, see your total score and performance rating

### Example Session
```
🎯 Welcome to the Multiple-Choice Quiz Game!
===========================================

📚 Available Categories:
----------------------
1. Programming (3 questions)
2. Science (2 questions)
3. History (2 questions)

Choose a category (enter number): 1

🚀 Starting Programming Quiz!
======================

Question 1/3:
❓ What does 'GOROOT' environment variable point to in Go?

  1. Go source code location
  2. Go binary installation directory
  3. Go workspace directory
  4. Go module cache

Enter your answer (1-4): 2
✅ Correct!
```

### Adding New Questions

To add new questions or categories, modify the `initializeQuestions()` method in `quiz_manager.go`:

```go
qm.categories["Your Category"] = []Question{
    {
        Text: "Your question?",
        Options: []string{"Option 1", "Option 2", "Option 3", "Option 4"},
        Answer: 0, // Index of correct answer (0-based)
    },
    // Add more questions...
}
```

### Building for Distribution

To create an executable binary:

```bash
go build -o quiz-game
./quiz-game
```

Enjoy the quiz game! 🎮