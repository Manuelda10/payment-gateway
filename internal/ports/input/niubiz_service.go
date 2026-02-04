package input

import "context"

type NiubizService interface {
	GetSecurityToken(ctx context.Context) (string, error)
}
