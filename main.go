package main

import (
    "log"
    "net/http"
		"os"
    "go-crud/internal/api"
		"go-crud/cmd"
)

func main() {
    http.HandleFunc("/questions", api.GetQuestions)
    http.HandleFunc("/submit", api.SubmitAnswers)

  
			// Check if the user has provided CLI arguments (e.g., "start" for the quiz)
			if len(os.Args) > 1 {
					// If there are CLI arguments, run the CLI
					cmd.Execute()
			} else {
					// Otherwise, start the server
					log.Println("Starting server on :8080...")
					log.Fatal(http.ListenAndServe(":8080", nil))
			}
}