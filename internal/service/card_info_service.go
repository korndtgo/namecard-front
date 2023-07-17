package service

import (
	"card-service/internal/config"
	"card-service/internal/entity"
	"card-service/internal/model"
	"card-service/internal/repository"
	"card-service/internal/util/log"
	"context"
	"github.com/emersion/go-vcard"
	"os"
)

const (
	LANGUAGE_TH = "TH"
	LANGUAGE_EN = "EN"
)

type CardInfoService struct {
	Repository repository.RepositoryGateway
	config     *config.Configuration
	logger     *log.Logger
}

func NewCardInfoService(
	repoGateway repository.RepositoryGateway,
	logger *log.Logger,
	config *config.Configuration,
) ICardInfoService {
	return &CardInfoService{
		Repository: repoGateway,
		logger:     logger,
		config:     config,
	}
}

type ICardInfoService interface {
	GetById(ctx context.Context, req model.CardRequest) (*entity.Card, error)
	GetByIdTh(ctx context.Context, req model.CardRequest) (string, error)
	GetByIdEn(ctx context.Context, req model.CardRequest) (string, error)
	CreateVCard(ctx context.Context, language string, req model.CardDto) (*os.File, error)
}

func (c CardInfoService) GetById(ctx context.Context, req model.CardRequest) (*entity.Card, error) {
	card, err := c.Repository.CardRepository.GetById(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:GetById] Call Service error: ", err)
		return &entity.Card{}, err
	}

	return card, nil
}

func (c CardInfoService) GetByIdTh(ctx context.Context, req model.CardRequest) (string, error) {
	card, err := c.Repository.CardRepository.GetById(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:GetByIdTh] Call Service error: ", err)
		return "", err
	}

	_, err = c.CreateVCard(ctx, LANGUAGE_TH, model.CardDto{
		Id:             card.Id,
		Uuid:           card.Uuid,
		EmployeeId:     card.EmployeeId,
		NameTh:         card.NameTh,
		NameEn:         card.NameEn,
		SurnameTh:      card.SurnameTh,
		SurnameEn:      card.SurnameEn,
		NicknameTh:     card.NicknameTh,
		NicknameEn:     card.NicknameEn,
		Email:          card.Email,
		ContactNumber1: card.ContactNumber1,
		ContactNumber2: card.ContactNumber2,
		LineId:         card.LineId,
		PositionTh:     card.PositionTh,
		PositionEn:     card.PositionEn,
		DepartmentTh:   card.DepartmentTh,
		DepartmentEn:   card.DepartmentEn,
		CompanyId:      card.CompanyId,
		Company:        model.CompanyDto{},
		CreatedAt:      card.CreatedAt,
		UpdatedAt:      card.UpdatedAt,
		DeletedAt:      card.DeletedAt,
		DownloadUrlTh:  "",
		DownloadUrlEn:  "",
	})

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:GetByIdTh] Call Create vCard error: ", err)
		return "", err
	}

	return "", nil
}

func (c CardInfoService) GetByIdEn(ctx context.Context, req model.CardRequest) (string, error) {
	card, err := c.Repository.CardRepository.GetById(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:GetByIdEn] Call Service error: ", err)
		return "", err
	}

	_, err = c.CreateVCard(ctx, LANGUAGE_EN, model.CardDto{
		Id:             card.Id,
		Uuid:           card.Uuid,
		EmployeeId:     card.EmployeeId,
		NameTh:         card.NameTh,
		NameEn:         card.NameEn,
		SurnameTh:      card.SurnameTh,
		SurnameEn:      card.SurnameEn,
		NicknameTh:     card.NicknameTh,
		NicknameEn:     card.NicknameEn,
		Email:          card.Email,
		ContactNumber1: card.ContactNumber1,
		ContactNumber2: card.ContactNumber2,
		LineId:         card.LineId,
		PositionTh:     card.PositionTh,
		PositionEn:     card.PositionEn,
		DepartmentTh:   card.DepartmentTh,
		DepartmentEn:   card.DepartmentEn,
		CompanyId:      card.CompanyId,
		Company:        model.CompanyDto{},
		CreatedAt:      card.CreatedAt,
		UpdatedAt:      card.UpdatedAt,
		DeletedAt:      card.DeletedAt,
		DownloadUrlTh:  "",
		DownloadUrlEn:  "",
	})

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:GetByIdEn] Call Create vCard error: ", err)
		return "", err
	}

	return "", nil
}

func (c CardInfoService) CreateVCard(ctx context.Context, language string, req model.CardDto) (*os.File, error) {
	destFile, err := os.Create("cards.vcf")
	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:CreateVCard]  error ", err)
		return nil, err
	}
	defer destFile.Close()

	var (
		// card is a map of strings to []*vcard.Field objects
		card vcard.Card

		// destination where the vcard will be encoded to
		enc = vcard.NewEncoder(destFile)
	)

	// set only the value of a field by using card.SetValue.
	// This does not set parameters
	if language == LANGUAGE_TH {
		card.SetValue(vcard.FieldName, req.NameTh)
		card.SetValue(vcard.FieldTitle, req.PositionTh)
		card.SetValue(vcard.FieldOrganization, req.Company.NameTh)
	} else {
		card.SetValue(vcard.FieldName, req.NameEn)
		card.SetValue(vcard.FieldTitle, req.PositionEn)
		card.SetValue(vcard.FieldOrganization, req.Company.NameEn)
	}
	card.SetValue(vcard.FieldTelephone, req.ContactNumber1)
	card.SetValue(vcard.FieldEmail, req.Email)

	// make the vCard version 4 compliant
	vcard.ToV4(card)
	err = enc.Encode(card)
	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[Service:CreateVCard] Encode  error ", err)
		return nil, err
	}

	return destFile, nil
}
