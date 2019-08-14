package debug

import "context"

type ClientRepository interface {
	Config(context.Context) (ClientWithCtxer, error)
	Connect(ClientWithCtxer) error
}
