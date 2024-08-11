package middleware

import (
	"chat-app/helper"
	"context"
	"math/rand"
	"net/http"

	"golang.org/x/exp/slog"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, helper.RequestId, rand.Intn(10000))
		r = r.WithContext(ctx)
		slog.Info("logging", "requestId", ctx.Value(helper.RequestId))
		next(w, r)
	}
}
