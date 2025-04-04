package helper

import (
	"context"
)

var PsqlKey = PsqlContextKey("psql")

type PsqlContextKey string

func FindContext(ctx context.Context, k PsqlContextKey) any {
	if v := ctx.Value(k); v != nil {
		return v
	}

	return nil
}
