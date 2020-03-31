package cluster

import "context"

type tokenAuth struct {
	token map[string]string
}

func (ta *tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return ta.token, nil
}

func (ta *tokenAuth) RequireTransportSecurity() bool {
	return false
}
