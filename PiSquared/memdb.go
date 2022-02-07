package PiSquared

// status contains current bot status for the user.
type status int

const (
	startMessage status = iota
	subjectSelected
	waitingResponseFromUser
)

// subject contains current subject selected by the user.
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

// user contains current status of the user saved in memory
// and the answer expected if a question is pending.
// It's used by bot to answer appropriately to the user.
type user struct {
	s                status
	sub              subject
	lastQuizQuestion string
	rightAnswer      string
}
