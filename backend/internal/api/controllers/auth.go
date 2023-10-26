package controllers

import "github.com/raphael-foliveira/fiber-react/backend/internal/api/services"

type auth struct {
	service services.Auth
}

func NewAuth(service services.Auth) *auth {
	return &auth{service}
}
