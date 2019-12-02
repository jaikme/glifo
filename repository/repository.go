package repository

import (
	"context"
	"glifo/models"
)

// SessionRepository explain...
type SessionRepository interface {
	Create(ctx context.Context, p *models.Session) (int64, error)
}
