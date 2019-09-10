package event

import "context"

type Repository interface {
	Load(ctx context.Context, prefix string) (<-chan []byte, error)
	Save(ctx context.Context, key string, payload []byte) error
}
