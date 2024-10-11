package generate

import "context"

type Generator interface {
	Generate(ctx context.Context) (string, error)
}
