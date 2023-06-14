// Copyright (c) 2022 - 2023 Erik Pellizzon
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package src

import (
	"encoding/json"
	"math/rand"
)

// Question represents a single question.
type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// Questions struct contains all the questions available.
type Questions struct {
	ComputerScience []Question `json:"computerScience"`
	Geography       []Question `json:"geography"`
	History         []Question `json:"history"`
}

var questions Questions

// LoadQuestions loads all the questions from json data.
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

// Attribute a percentage score to the User answer.
// From 0 to 100%.
func evalAnswer(rightAnswer, userAnswer string) float32 {
	value := dotProduct(vector(rightAnswer), vector(userAnswer))
	if value > 1 {
		value = 1
	}
	return value * 100
}
