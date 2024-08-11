package controller

import (
	"context"
	"fmt"
	"net/http"
)

func Chat(context context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "chat started")
	return nil
}
