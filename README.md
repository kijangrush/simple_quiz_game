## Go Quiz Game

A simple command-line multiple-choice quiz game written in Go with scoring and topic categories.

### Features

- 🎯 Multiple topic categories (Programming Basics, Go Language, General Knowledge)
- 📊 Real-time scoring and performance feedback
- 🔄 Play multiple rounds with different categories
- 🎮 User-friendly command-line interface
- ✅ Immediate feedback on answers

### How to Use

1. **Prerequisites**: Make sure you have Go installed on your system (version 1.16 or higher).

2. **Setup**:
   ```bash
   # Clone or create the project directory
   mkdir quiz-game
   cd quiz-game
   
   # Create the files above or copy them into the directory
   # The project structure should be:
   # quiz-game/
   # ├── main.go
   # ├── quiz.go
   # └── go.mod
   ```

3. **Run the game**:
   ```bash
   go run .
   ```

   Or build and run:
   ```bash
   go build
   ./quiz-game
   ```

### Game Instructions

1. When you start the game, you'll see a list of available categories
2. Choose a category by entering the corresponding number
3. Answer each multiple-choice question by entering the number of your choice
4. After completing a quiz, you'll see your score and performance feedback
5. You can choose to play again with the same or different categories

### Categories Included

- **Programming Basics**: Fundamental computer science concepts
- **Go Language**: Specific questions about the Go programming language
- **General Knowledge**: Miscellaneous trivia questions

### Customization

You can easily add more categories or questions by modifying the `GetCategories()` function in `quiz.go`. Simply add new `Category` structs with your desired questions.

### Example Game Session

```
🎯 Welcome to the Go Quiz Game!
================================

Available Categories:
1. Programming Basics
2. Go Language
3. General Knowledge
4. Exit

Choose a category (1-4): 1

🏁 Starting Programming Basics Quiz!
========================================

Question 1/3:
❓ What does CPU stand for?
   1. Computer Processing Unit
   2. Central Processing Unit
   3. Central Program Utility
   4. Computer Program Unit

Enter your answer (1-4): 2
✅ Correct!
```

Enjoy playing the quiz game! 🎮