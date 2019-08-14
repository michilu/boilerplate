// Created by interfacer; DO NOT EDIT

package debug

import (
	"github.com/michilu/boilerplate/service/debug"
)

// Configer is an interface generated for "github.com/michilu/boilerplate/application/debug.Config".
type Configer interface {
	Config(struct{}) (debug.ClientWithCtxer, error)
	Connect(debug.ClientWithCtxer) (struct{}, error)
}
