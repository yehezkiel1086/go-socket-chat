package port

import (
	"context"
	"go-socket/internal/core/domain"
)

type ClientService interface {
	JoinRoom(ctx context.Context, cl *domain.Client)
}
