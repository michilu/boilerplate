// Created by interfacer; DO NOT EDIT

package debug

import (
	"context"
	"github.com/michilu/boilerplate/service/debug"
)

// Configer is an interface generated for "github.com/michilu/boilerplate/application/debug.Config".
type Configer interface {
	Config(context.Context) (debug.ClientWithContexter, error)
	Connect(debug.ClientWithContexter) (context.Context, error)
}
