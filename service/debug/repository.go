package debug

import "context"

type ClientRepository interface {
	Config(context.Context) (ClientWithContexter, error)
	Connect(ClientWithContexter) error
}
