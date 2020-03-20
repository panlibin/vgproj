package igame

type INameManager interface {
	GrabPlayerName(name string) bool
	ReleasePlayerName(name string)
	GrabGuildName(name string) bool
	ReleaseGuildName(name string)
}
