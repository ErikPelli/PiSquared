package PiSquared

import (
	"encoding/json"
	"math/rand"
)

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Questions struct {
	ComputerScience []Question `json:"computerScience"`
	Geography       []Question `json:"geography"`
	History         []Question `json:"history"`
}

var questions Questions

func LoadQuestions(jsonData []byte) error {
	return json.Unmarshal(jsonData, &questions)
}

// getQuestions returns a new question.
// First returned variable is the question and the second is the answer.
func getQuestion(s subject) Question {
	var q []Question
	switch s {
	case computerScience:
		q = questions.ComputerScience
	case geography:
		q = questions.Geography
	case history:
		q = questions.History
	default:
		return Question{}
	}
	return q[rand.Intn(len(q))]
}

// Attribute a percentage score to the answer.
// From 0 to 100%.
func evalAnswer(question, answer string) float32 {
	// TODO
	return 50
}
