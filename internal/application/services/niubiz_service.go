package services

import (
	"context"
	"fmt"
	"payment-gateway/internal/ports/input"
	"payment-gateway/internal/ports/output"
)

type niubizService struct {
	log     output.Logger
	gateway output.NiubizGateway
}

func NewNiubizService(log output.Logger, gateway output.NiubizGateway) input.NiubizService {
	return &niubizService{log: log, gateway: gateway}
}

func (s *niubizService) GetSecurityToken(ctx context.Context) (string, error) {
	s.log.Info(ctx, "niubiz: getting security token")

	token, err := s.gateway.GetSecurityToken(ctx)
	if err != nil {
		s.log.Error(ctx, "niubiz: security token failed", err)
		return "", err
	}

	if token == "" {
		err := fmt.Errorf("niubiz security token is empty")
		s.log.Error(ctx, "niubiz: security token empty", err)
		return "", err
	}

	s.log.Info(ctx, "niubiz: security token obtained successfully")
	return token, nil
}
