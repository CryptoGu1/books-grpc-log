package repo

import (
	"context"
	audit "github.com/CryptoGu1/books-grpc-log/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Database
}

func NewAudit(db *mongo.Database) *Audit {
	return &Audit{db: db}
}

func (r *Audit) Insert(ctx context.Context, item audit.)
