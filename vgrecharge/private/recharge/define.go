package recharge

const (
	platformApple      = 1
	platformGooglePlay = 2
)

type sdkParam struct {
	pfId  int32
	name  string
	appId string
	keys  []string
}
