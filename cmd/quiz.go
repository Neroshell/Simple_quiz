package cmd

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "sort"
)

// QuizQuestion struct to hold question and answers
type QuizQuestion struct {
    Question string
    Answers  []string
    Correct  int // Index of the correct answer
}

// Questions slice to store the quiz questions
var Questions = []QuizQuestion{
    {
        Question: "What is the largest mammal in the world?",
        Answers:  []string{"Elephant", "Blue Whale", "Great White Shark", "Giraffe"},
        Correct:  1, // Blue Whale is the correct answer
    },
    {
        Question: "Which country is known as the Land of the Rising Sun?",
        Answers:  []string{"China", "Japan", "Thailand", "South Korea"},
        Correct:  1, // Japan is the correct answer
    },
    {
        Question: "What is the hardest natural substance on Earth?",
        Answers:  []string{"Gold", "Diamond", "Iron", "Quartz"},
        Correct:  1, // Diamond is the correct answer
    },
    {
        Question: "What is the main ingredient in guacamole?",
        Answers:  []string{"Tomato", "Avocado", "Pepper", "Onion"},
        Correct:  1, // Avocado is the correct answer
    },
    {
        Question: "What is the capital city of Canada?",
        Answers:  []string{"Toronto", "Vancouver", "Ottawa", "Montreal"},
        Correct:  2, // Ottawa is the correct answer
    },
    {
        Question: "In which year did the Titanic sink?",
        Answers:  []string{"1912", "1914", "1916", "1918"},
        Correct:  0, // 1912 is the correct answer
    },
    {
        Question: "Which planet is known for its rings?",
        Answers:  []string{"Earth", "Mars", "Jupiter", "Saturn"},
        Correct:  3, // Saturn is the correct answer
    },
    {
        Question: "What is the smallest prime number?",
        Answers:  []string{"1", "2", "3", "5"},
        Correct:  1, // 2 is the correct answer
    },
    {
        Question: "Which element has the atomic number 1?",
        Answers:  []string{"Oxygen", "Hydrogen", "Helium", "Carbon"},
        Correct:  1, // Hydrogen is the correct answer
    },
    {
        Question: "Who painted the Mona Lisa?",
        Answers:  []string{"Vincent van Gogh", "Pablo Picasso", "Leonardo da Vinci", "Claude Monet"},
        Correct:  2, // Leonardo da Vinci is the correct answer
    },
}

// Slice to store all quiz scores
var allScores []int

// LoadScores reads scores from a file
func LoadScores() {
    file, err := ioutil.ReadFile("scores.json")
    if err != nil {
        fmt.Println("No previous scores found.")
        return
    }
    err = json.Unmarshal(file, &allScores)
    if err != nil {
        fmt.Println("Error loading scores:", err)
    }
}

// SaveScores writes scores to a file
func SaveScores() {
    file, err := json.MarshalIndent(allScores, "", " ")
    if err != nil {
        fmt.Println("Error saving scores:", err)
        return
    }
    ioutil.WriteFile("scores.json", file, 0644)
}

// StartQuiz function with persistent score storage
func StartQuiz() {
    // Load previous scores from file
    LoadScores()

    correctAnswers := 0
    for _, q := range Questions {
        fmt.Println(q.Question)
        for i, ans := range q.Answers {
            fmt.Printf("%d: %s\n", i+1, ans)
        }

        var answer int
        fmt.Print("Select your answer (1-4): ")
        fmt.Scan(&answer)

        if answer-1 == q.Correct {
            fmt.Println("Correct!")
            correctAnswers++
        } else {
            fmt.Println("Incorrect. The correct answer was:", q.Answers[q.Correct])
        }
    }

    fmt.Printf("You got %d out of %d correct.\n", correctAnswers, len(Questions))

    // Add the current user's score
    allScores = append(allScores, correctAnswers)

    // Save the updated scores to file
    SaveScores()

    // Compare with others
    compareScore(correctAnswers)
}

// compareScore function to compare the user's score with others
func compareScore(currentScore int) {
    sort.Ints(allScores)
    rank := 0
    for i, score := range allScores {
        if score == currentScore {
            rank = i + 1
            break
        }
    }
    totalUsers := len(allScores)
    percentile := float64(rank-1) / float64(totalUsers) * 100
    fmt.Printf("You performed better than %.2f%% of all quizzers.\n", percentile)
}
