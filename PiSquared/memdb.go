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
	math subject = iota
	computerScience
	geography
	history
)

// user contains current status of the user saved in memory
// and the answer expected if a question is pending.
// It's used by bot to answer appropriately to the user.
type user struct {
	s                status
	sub              subject
	lastQuizQuestion string
	rightAnswer      string
}
