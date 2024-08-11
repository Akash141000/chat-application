package controller

import (
	"chat-app/helper"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(context context.Context, w http.ResponseWriter, r *http.Request) error {
	loginReq := &LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		return fmt.Errorf("invalid parameters")
	}
	// validate struct
	if err := helper.Validate.Struct(loginReq); err != nil {
		return err
	}

	// var authToken string
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "chat-app",
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := unsignedToken.SignedString([]byte(helper.SigningKey))
	if err != nil {
		return err
	}

	resp := &LoginResponse{
		Token: signedToken,
	}

	json.NewEncoder(w).Encode(resp)

	return nil
}

func SignUp(context context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "signup")
	return nil
}
