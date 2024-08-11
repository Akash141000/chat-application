package helper

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const SigningKey = "testingkey"

type UserId string

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Response)

type ControllerFunc func(context.Context, http.ResponseWriter, *http.Request) error

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type CtxRequestId string

const RequestId CtxRequestId = "requestId"
