package constants

type Actions int

const (
	DISCARD           Actions = 0
	ALLOW             Actions = 1
	REFUSE            Actions = 2
	DISCARD_PASSIVE   Actions = 4
	TERMINATE_PASSIVE Actions = 8
	TERMINATE         Actions = 9
	TERMINATE_FAILED  Actions = 10
	PERMIT            Actions = 11
	TERMINATE_RESET   Actions = 12
	BLOCK             Actions = 13
)
