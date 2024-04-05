package rest

import (
	"trec/service"

	"go.uber.org/zap"
)

// Controller is the interface for all controllers.
type RestController interface {
	Auth() AuthController
}

// controllerImpl is the implementation of Controller.
type restControllerImpl struct {
	authController AuthController
}

// NewRestController creates a new Controller.
func NewRestController(service service.Service, clientRedirectURL string, logger *zap.Logger) RestController {
	return restControllerImpl{
		authController: NewAuthController(service.Auth(), clientRedirectURL, logger),
	}
}

// Auth returns the AuthController.
func (i restControllerImpl) Auth() AuthController {
	return i.authController
}
