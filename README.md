# Go Quiz Game

A simple command-line multiple-choice quiz game written in Go with scoring and topic categories.

## Features

- 🎯 Multiple quiz topics (Go Programming, Computer Science, General Knowledge)
- 📊 Score tracking and percentage calculation
- 🏆 Performance feedback based on score
- 🎮 Interactive command-line interface

## Prerequisites

- Go 1.16 or higher installed on your system

## Installation & Usage

1. **Create a new directory for the project:**
   ```bash
   mkdir go-quiz-game
   cd go-quiz-game
   ```

2. **Create the files:**
   - Copy the code above into three separate files:
     - `main.go`
     - `quiz.go`
     - `go.mod`

3. **Run the game:**
   ```bash
   go run main.go quiz.go
   ```

## How to Play

1. **Start the game:** Run the program using `go run main.go quiz.go`

2. **Choose a topic:** Select from the available quiz categories by entering the corresponding number

3. **Answer questions:** For each question, read the prompt and options, then enter the number of your chosen answer

4. **View results:** After completing all questions, see your score and performance feedback

## Game Structure

- **Topics:** Currently includes three categories:
  - Go Programming Basics
  - Computer Science Fundamentals
  - General Knowledge

- **Scoring:** Each correct answer adds 1 point to your score
- **Performance:** Get feedback based on your percentage score

## Customization

You can easily add more questions or topics by modifying the `LoadQuizzes()` function in `quiz.go`. Each quiz topic contains:

- `Name`: The display name of the topic
- `Questions`: Array of question objects with:
  - `Text`: The question prompt
  - `Options`: Array of possible answers
  - `Answer`: Index (0-based) of the correct answer

## Example Game Session

```
🎯 Welcome to the Go Quiz Game!
===============================

Available Quiz Topics:
1. Go Programming Basics
2. Computer Science Fundamentals
3. General Knowledge

Select a topic (number): 1

📝 Starting Quiz: Go Programming Basics
========================

Question 1/3: What keyword is used to declare a variable in Go?
  1. var
  2. let
  3. const
  4. variable
Your answer (number): 1
✅ Correct!

... more questions ...

🎉 Quiz Completed!
Your score: 2/3 (66.7%)
💪 Keep practicing!
```

Enjoy testing your knowledge with this simple Go quiz game!