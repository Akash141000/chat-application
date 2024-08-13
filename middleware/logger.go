package middleware

import (
	"chat-app/helper"
	"context"
	"math/rand"
	"net/http"
	"strconv"

	"golang.org/x/exp/slog"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, helper.RequestId, strconv.Itoa(rand.Intn(10000)))
		r = r.WithContext(ctx)
		slog.Info("logging", "requestId", ctx.Value(helper.RequestId))
		//testing savelog
		helper.SaveLog("reqId", ctx.Value(helper.RequestId).(string))
		next(w, r)
	}
}
