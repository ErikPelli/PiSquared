package PiSquared

import cmap "github.com/orcaman/concurrent-map"

// status contains current bot status for the user.
type status int

const (
	startMessage status = iota
	subjectSelected
	waitingResponseFromUser
)

// user contains current status of the user saved in memory
// and the answer expected if a question is pending.
// It's used by bot to answer appropriately to the user.
type user struct {
	s                status
	lastQuizQuestion string
	rightAnswer      string
}

// memory is the association between user chat id (key) and its status (value).
type memory cmap.ConcurrentMap
