package service

import (
	"go.uber.org/dig"
)

// Service ...
type Service struct {
	ServiceGateway ServiceGateway
}

// ServiceGateway ...
type ServiceGateway struct {
	dig.In

	CardService     ICardService
	CardInfoService ICardInfoService

	//ClaimService           IClaimService
}

// NewService ...
func NewService(sg ServiceGateway) *Service {
	return &Service{
		ServiceGateway: sg,
	}
}
