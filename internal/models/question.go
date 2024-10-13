package models

type QuestionAPI struct {
	Question        string   `json:"question"`
	CorrectAnswer   string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type UserScore struct {
	UserID   string
	Accuracy float64
}


var UserScores = []UserScore{}
