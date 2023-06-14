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
	"gorm.io/gorm"
	"time"
)

// status contains current bot status for the User.
type status int

const (
	startMessage status = iota
	subjectSelected
	waitingResponseFromUser
)

// subject contains current subject selected by the User.
type subject int

const (
	computerScience subject = iota
	geography
	history
)

// String returns string representation of current subject.
func (s subject) String() string {
	switch s {
	case computerScience:
		return "Computer Science"
	case geography:
		return "Geography"
	case history:
		return "History"
	default:
		return ""
	}
}

// User contains current status of the User and the
// answer expected if a question is pending.
// It's used by bot to answer appropriately to the User.
type User struct {
	ChatID           int64   `gorm:"primaryKey"`
	S                status  `gorm:"type:int"`
	Sub              subject `gorm:"type:int"`
	LastQuizQuestion string
	RightAnswer      string

	// GORM fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
