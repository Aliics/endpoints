package endpoints

import (
	"context"
	"github.com/google/uuid"
)

type Context struct {
	context.Context

	AccountID *uuid.UUID
}

func BackgroundContext() Context {
	return Context{context.Background(), nil}
}

func FromContext(ctx context.Context) Context {
	var accountID *uuid.UUID

	if value, ok := ctx.Value("AccountID").(uuid.UUID); ok {
		accountID = &value
	}

	return Context{ctx, accountID}
}
