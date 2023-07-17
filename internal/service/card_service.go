package service

import (
	"card-service/internal/config"
	"card-service/internal/entity"
	"card-service/internal/model"
	"card-service/internal/repository"
	"card-service/internal/util/log"
	"context"
)

type ICardService interface {
	FindAll(ctx context.Context, req model.QueryCard) ([]entity.Card, error)
	Create(ctx context.Context, req model.CreateCardDto) (*entity.Card, error)
	Update(ctx context.Context, req model.UpdateCardDto) (*entity.Card, error)
	Delete(ctx context.Context, req model.CardRequest) error
}

type CardService struct {
	Repository repository.RepositoryGateway
	config     *config.Configuration
	logger     *log.Logger
}

func (c CardService) FindAll(ctx context.Context, req model.QueryCard) ([]entity.Card, error) {
	list, _, _, err := c.Repository.CardRepository.GetAll(ctx, req)
	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:GetCompanyIdFromUserId] Call Repo GetCompanyIdFromUserId error: ", err)
		return []entity.Card{}, err
	}

	return list, nil
}

func (c CardService) Create(ctx context.Context, req model.CreateCardDto) (*entity.Card, error) {
	res, err := c.Repository.CardRepository.Create(ctx, req)
	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:Create] Call Repo Create error: ", err)
		return &entity.Card{}, err
	}

	return res, nil
}

func (c CardService) Update(ctx context.Context, req model.UpdateCardDto) (*entity.Card, error) {
	res, err := c.Repository.CardRepository.Update(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:Update] Call Repo Update error: ", err)
		return &entity.Card{}, err
	}

	return res, nil
}

func (c CardService) Delete(ctx context.Context, req model.CardRequest) error {
	//soft delete company
	err := c.Repository.CardRepository.Delete(ctx, req)

	return err
}

func NewCardService(
	repoGateway repository.RepositoryGateway,
	logger *log.Logger,
	config *config.Configuration,
) ICardService {
	return &CardService{
		Repository: repoGateway,
		logger:     logger,
		config:     config,
	}
}
