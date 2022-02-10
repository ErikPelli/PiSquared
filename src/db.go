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
