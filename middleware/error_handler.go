package middleware

import (
	"chat-app/helper"
	"context"
	"fmt"
	"net/http"
	"time"
)

func ErrorHandler(controller helper.ControllerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check the method
		if r.Method != method {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "url not found")
			return
		}

		//update the timeout time later from 3
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
		defer cancel()
		if err := controller(ctx, w, r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}
