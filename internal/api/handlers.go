package api



import (
	"encoding/json"
	"fmt"
	"net/http"
	"html"
	"go-crud/internal/models"
	"math/rand"
	"time"
	"strconv"
)

type TriviaResponse struct {
	ResponseCode int                 `json:"response_code"`
	Results      []models.QuestionAPI `json:"results"`
}


// Fetches questions from the Open Trivia API
func FetchQuestionsFromAPI(amount int) ([]models.QuestionAPI, error) {
	url := fmt.Sprintf("https://opentdb.com/api.php?amount=%d&type=multiple", amount)

	res, err := http.Get(url)
	if err != nil {
			return nil, fmt.Errorf("error fetching questions: %v", err)
	}
	defer res.Body.Close()

	var trivia TriviaResponse
	if err := json.NewDecoder(res.Body).Decode(&trivia); err != nil {
			return nil, fmt.Errorf("error decoding trivia response: %v", err)
	}

	return trivia.Results, nil
}

// GetQuestions serves questions to the CLI, allowing dynamic number of questions
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil || amount < 1 {
			http.Error(w, "Invalid amount parameter", http.StatusBadRequest)
			return
	}

	questions, err := FetchQuestionsFromAPI(amount)
	if err != nil {
			http.Error(w, "Failed to fetch questions", http.StatusInternalServerError)
			return
	}

	// Escape HTML entities in questions and answers
	for i, q := range questions {
			questions[i].Question = html.UnescapeString(q.Question)
			questions[i].CorrectAnswer = html.UnescapeString(q.CorrectAnswer)
			for j, a := range q.IncorrectAnswers {
					questions[i].IncorrectAnswers[j] = html.UnescapeString(a)
			}
	}

	// Shuffle the questions
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	randGen.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
	})

	json.NewEncoder(w).Encode(questions)
}

// SubmitAnswers processes the user's answers and calculates their score
func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
			User      string                     `json:"user"`
			Answers   map[int]string             `json:"answers"`
			Questions []models.QuestionAPI       `json:"questions"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
	}

	correct := 0
	// Evaluate the user's answers
	for i, userAnswer := range requestData.Answers {
			if userAnswer == requestData.Questions[i].CorrectAnswer {
					correct++
			}
	}

	userID := requestData.User
	result := submitResults(userID, correct)

	json.NewEncoder(w).Encode(result)
}

func submitResults(userID string, score int) string {
	// Check if the user already exists and update the score
	found := false
	for i, userScore := range models.UserScores {
			if userScore.UserID == userID {
					models.UserScores[i].Score = score  // Overwrite the existing score
					found = true
					break
			}
	}

	if !found {
			models.UserScores = append(models.UserScores, models.UserScore{UserID: userID, Score: score})
	}

	// Compare with other users
	totalUsers := len(models.UserScores)
	betterThanCount := 0

	for _, userScore := range models.UserScores {
			if userScore.Score < score {
					betterThanCount++
			}
	}

	percentile := (float64(betterThanCount) / float64(totalUsers)) * 100
	return fmt.Sprintf("You scored %d. You were better than %.2f%% of all quizzers.", score, percentile)
}
