package rest

import (
	"trec/config"
	"trec/service"

	"go.uber.org/zap"
)

// Controller is the interface for all controllers.
type RestController interface {
	Auth() AuthController
	Talena() TalenaController
}

// controllerImpl is the implementation of Controller.
type restControllerImpl struct {
	authController   AuthController
	talenaController TalenaController
}

// NewRestController creates a new Controller.
func NewRestController(service service.Service, configs *config.Configurations, logger *zap.Logger) RestController {
	return restControllerImpl{
		authController:   NewAuthController(service.Auth(), configs.AzureADOAuth.ClientRedirectUrl, logger),
		talenaController: NewTalenaController(configs.Talena, logger),
	}
}

// Auth returns the AuthController.
func (i restControllerImpl) Auth() AuthController {
	return i.authController
}

// Talena returns the TalenaController.
func (i restControllerImpl) Talena() TalenaController {
	return i.talenaController
}
