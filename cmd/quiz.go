package cmd

import (
    "fmt"
    "net/http"
    "encoding/json"
    "bufio"
    "os"
    "strings"
    "bytes"
    "strconv"
		"math/rand"

    "github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
    Use:   "start",
    Short: "Start the quiz",
    Run: func(cmd *cobra.Command, args []string) {
        takeQuiz()
    },
}

func takeQuiz() {
	reader := bufio.NewReader(os.Stdin)

	for {
			// Step 1: Ask how many questions the user wants
			fmt.Print("How many questions do you want to answer? ")
			questionCountInput, _ := reader.ReadString('\n')
			questionCountInput = strings.TrimSpace(questionCountInput)
			questionCount, err := strconv.Atoi(questionCountInput)
			if err != nil || questionCount < 1 {
					fmt.Println("Invalid number of questions. Please enter a valid number.")
					continue
			}

			// Step 2: Ask for the user's name
			fmt.Print("Enter your name: ")
			userNameInput, _ := reader.ReadString('\n')
			userNameInput = strings.TrimSpace(userNameInput)

			// Ensure the name is not empty
			if userNameInput == "" {
					fmt.Println("Name cannot be empty. Please enter your name.")
					continue
			}
			userName := userNameInput

			// Fetch questions from the local REST API, using the question count
			questionURL := fmt.Sprintf("http://localhost:8080/questions?amount=%d", questionCount)
			res, err := http.Get(questionURL)
			if err != nil {
					fmt.Println("Failed to get questions:", err)
					return
			}
			defer res.Body.Close()

			var questions []map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&questions); err != nil {
					fmt.Println("Failed to decode response:", err)
					return
			}

			userAnswers := make(map[int]string)

			// Step 3: Play the quiz
			for i, question := range questions {
					fmt.Printf("Question %d: %s\n", i+1, question["question"])
					answers := append(question["incorrect_answers"].([]interface{}), question["correct_answer"])

					// Shuffle the answers
					rand.Shuffle(len(answers), func(i, j int) {
							answers[i], answers[j] = answers[j], answers[i]
					})

					for idx, answer := range answers {
							fmt.Printf("%d: %s\n", idx+1, answer)
					}

					// Get the user's answer
					response, _ := reader.ReadString('\n')
					response = strings.TrimSpace(response)
					selected, err := strconv.Atoi(response)
					if err != nil || selected < 1 || selected > len(answers) {
							fmt.Println("Invalid answer.")
							continue
					}

					// Store the user's selected answer (as a string)
					userAnswers[i] = answers[selected-1].(string)
			}

			// Step 4: Submit answers to the server, along with the user's name and questions
			postAnswers(userName, userAnswers, questions)

			// Step 5: Ask if they want to play again
			fmt.Print("Do you want to play again? (yes/no): ")
			playAgainInput, _ := reader.ReadString('\n')
			playAgain := strings.TrimSpace(strings.ToLower(playAgainInput))

			if playAgain == "no" {
					fmt.Println("Thank you for playing!")
					return // Exit the program
			}
	}
}


func postAnswers(userName string, answers map[int]string, questions []map[string]interface{}) {
	data := map[string]interface{}{
			"user":     userName,
			"answers":  answers,
			"questions": questions,
	}

	jsonData, _ := json.Marshal(data)
	res, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
			fmt.Println("Failed to submit answers:", err)
			return
	}
	defer res.Body.Close()

	var result string
	json.NewDecoder(res.Body).Decode(&result)
	fmt.Println(result)
}

func init() {
    rootCmd.AddCommand(startCmd)
}
