## How to Use the Quiz Game

### Prerequisites
- Go 1.16 or higher installed on your system

### Installation & Running

1. **Create a new directory for the project:**
   ```bash
   mkdir quiz-game
   cd quiz-game
   ```

2. **Create the files:**
   Copy the code from each section above into their respective files:
   - `main.go`
   - `models.go` 
   - `quiz.go`
   - `go.mod`

3. **Run the game:**
   ```bash
   go run .
   ```
   Or build and run:
   ```bash
   go build
   ./quiz-game
   ```

### Game Features

- **Multiple Categories**: Choose from 3 different quiz categories
  - Programming Basics
  - Go Language  
  - General Knowledge

- **Scoring System**: Tracks your score and calculates percentage
- **User-friendly Interface**: Clear instructions and feedback
- **Input Validation**: Handles invalid inputs gracefully

### How to Play

1. Run the program
2. Select a category by entering 1, 2, or 3
3. Answer each multiple-choice question by typing A, B, C, or D
4. View your final score and performance feedback at the end

### Customization

To add more questions or categories, modify the `getCategories()` function in `models.go`. Each category contains:
- `Name`: The display name of the category
- `Questions`: Array of Question objects with:
  - `Text`: The question text
  - `Options`: Array of 4 possible answers
  - `Answer`: Index (0-3) of the correct answer

### Example Game Session
```
🎯 Welcome to the Multiple-Choice Quiz Game!
===========================================

Available Categories:
1. Programming Basics
2. Go Language
3. General Knowledge

Select a category (1-3): 2

Starting Go Language Quiz...
========================

Question 1: Who created the Go programming language?
  A. Microsoft
  B. Google
  C. Apple
  D. Facebook
Your answer (A/B/C/D): B
✅ Correct!
...
```

Enjoy your quiz game! 🎮