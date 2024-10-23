## Multiple-Choice Quiz Game

A simple command-line multiple-choice quiz game written in Go with scoring and topic categories.

### Features

- 📚 Multiple topic categories
- ✅ Real-time scoring
- 🎯 Interactive command-line interface
- 📊 Final score display with percentage

### How to Use

1. **Prerequisites**: Make sure you have Go installed on your system.

2. **Setup**: 
   ```bash
   # Clone or create the project directory with all the files
   # The files should be: main.go, models.go, quiz.go, data.go, go.mod
   ```

3. **Run the game**:
   ```bash
   go run .
   ```
   Or build and run:
   ```bash
   go build -o quiz-game
   ./quiz-game
   ```

4. **Game Flow**:
   - The program will display available categories
   - Select a category by entering its number
   - Answer each question by entering the number of your choice
   - After completing all questions, your score will be displayed

### File Structure

- `main.go` - Entry point and main game flow
- `models.go` - Data structures (Question, Category)
- `quiz.go` - Quiz logic and scoring
- `data.go` - Sample quiz questions and categories
- `go.mod` - Go module definition

### Adding More Questions

To add more questions or categories, modify the `GetCategories()` function in `data.go`. Follow the existing structure:

```go
{
    Name: "Your Category Name",
    Questions: []Question{
        {
            Text: "Your question?",
            Options: []string{"Option 1", "Option 2", "Option 3", "Option 4"},
            Answer: 0, // Index of correct answer (0-based)
        },
        // Add more questions...
    },
}
```

### Example Output
```
🎯 Welcome to the Multiple-Choice Quiz Game!
===========================================

Available Categories:
1. Programming Fundamentals
2. Go Language
3. General Knowledge

Select a category (number): 2

Question 1/3:
❓ Which keyword is used to declare a variable in Go?
  1. var
  2. let
  3. const
  4. variable
Your answer (number): 1
✅ Correct!

🎊 Quiz Completed! 🎊
Your Score: 2/3 (66.7%)
```

Enjoy your quiz game! 🎮