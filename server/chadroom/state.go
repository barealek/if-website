package chadroom

type State string

const (
	StateWaitingUser State = "waiting:user"
	StateWaitingChad State = "waiting:chad"
	StateNotStarted  State = "not-started"
)
