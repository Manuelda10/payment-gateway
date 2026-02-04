package output

import "context"

type NiubizGateway interface {
	GetSecurityToken(ctx context.Context) (string, error)
}
