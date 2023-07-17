package repository

import (
	"card-service/internal/config"
	"card-service/internal/entity"
	"card-service/internal/infrastructure/database"
	"github.com/google/uuid"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"

	"card-service/internal/model"
	"card-service/internal/util/databasehelp"
	"card-service/internal/util/log"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type CardRepository struct {
	config         *config.Configuration
	db             *database.DB
	logger         *log.Logger
	TimeRepository ITimeRepository
	databaseHelp   *databasehelp.DatabaseHelp
}

type ICardRepository interface {
	GetAll(ctx context.Context, req model.QueryCard) ([]entity.Card, *string, *int64, error)
	GetById(ctx context.Context, req model.CardRequest) (*entity.Card, error)
	Create(ctx context.Context, req model.CreateCardDto) (*entity.Card, error)
	Update(ctx context.Context, req model.UpdateCardDto) (*entity.Card, error)
	Delete(ctx context.Context, req model.CardRequest) error
}

func NewCardRepository(
	config *config.Configuration,
	db *database.DB,
	logger *log.Logger,
) ICardRepository {
	return &CardRepository{
		config: config,
		logger: logger,
		db:     db,
	}
}

func (c CardRepository) GetById(ctx context.Context, req model.CardRequest) (*entity.Card, error) {
	c.logger.DebugWithID(ctx, log.AppLog, "[Repo:GetById] Called")

	transaction := c.db.Db.Begin()

	card := &entity.Card{}

	r := transaction.Where("id = ?", req.Id).Limit(1).Find(&card)

	if r.RowsAffected == 0 {
		err := errors.New("record not found")
		c.logger.ErrorWithID(ctx, log.AppLog, "[Repo:GetById data] GetById   error ", err)
		return nil, err
	}

	if err := transaction.Where("id = ?", req.Id).Model(&card).Updates(map[string]interface{}{
		"IsActive": false,
	}).Error; err != nil {
		transaction.Rollback()
		c.logger.ErrorWithID(ctx, log.AppLog, "[Repo:GetById data] GetById   error ", err)
		return nil, err
	}

	transaction.Commit()

	return card, nil
}

func (c CardRepository) GetAll(ctx context.Context, req model.QueryCard) ([]entity.Card, *string, *int64, error) {
	c.logger.DebugWithID(ctx, log.AppLog, "[Repository:GetCampaignList] Called")

	orm := c.db.Db.Debug().Model(&entity.Card{})

	var count *int64
	var tempCount int64
	count = &tempCount
	if err := orm.
		Session(&gorm.Session{}).
		Count(count).Error; err != nil {
		return nil, nil, nil, err
	}

	var results []entity.Card
	//eofCursor := "EOF"

	pageRules := make([]paginator.Rule, req.PageSize)
	sorting := paginator.DESC

	pageRules = append(pageRules, paginator.Rule{
		Key:   "Id",
		Order: sorting,
	})

	p := paginator.New(&paginator.Config{
		Rules: pageRules,
		Order: sorting,
		Limit: req.PageSize,
	})

	returnGorm, cur, err := p.Paginate(orm, &results)
	if err != nil {
		return nil, nil, nil, err
	}
	if returnGorm.Error != nil {
		return nil, nil, nil, returnGorm.Error
	}
	if cur.After == nil { // out of data
		return results, nil, count, nil
	}

	return results, cur.After, count, nil
}

func (c CardRepository) Create(ctx context.Context, req model.CreateCardDto) (*entity.Card, error) {
	c.logger.DebugWithID(ctx, log.AppLog, "[Repo:Create] Called")

	transaction := c.db.Db.Begin()

	company := &entity.Card{
		Uuid:           uuid.New().String(),
		EmployeeId:     req.EmployeeId,
		NameTh:         req.NameTh,
		NameEn:         req.NameEn,
		SurnameTh:      req.SurnameTh,
		SurnameEn:      req.SurnameEn,
		NicknameTh:     req.NicknameTh,
		NicknameEn:     req.NicknameEn,
		Email:          req.Email,
		ContactNumber1: req.ContactNumber1,
		ContactNumber2: req.ContactNumber2,
		LineId:         req.LineId,
		PositionTh:     req.PositionTh,
		PositionEn:     req.PositionEn,
		DepartmentTh:   req.DepartmentTh,
		DepartmentEn:   req.DepartmentEn,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := transaction.Create(&company).Error; err != nil {
		transaction.Rollback()
		c.logger.ErrorWithID(ctx, log.AppLog, "[Repo:Create] Create  error ", err)
		return nil, err
	}

	transaction.Commit()

	return company, nil
}

func (c CardRepository) Update(ctx context.Context, req model.UpdateCardDto) (*entity.Card, error) {
	c.logger.DebugWithID(ctx, log.AppLog, "[Repo:Update] Called")

	transaction := c.db.Db.Begin()

	card := &entity.Card{}

	r := transaction.Where("id = ?", req.Id).Limit(1).Find(&card)

	if r.RowsAffected == 0 {
		err := errors.New("record not found")
		c.logger.ErrorWithID(ctx, log.AppLog, "[Repo:Updated data] update   error ", err)
		return nil, err
	}

	card = &entity.Card{
		NameTh:         req.NameTh,
		NameEn:         req.NameEn,
		SurnameTh:      req.SurnameTh,
		SurnameEn:      req.SurnameEn,
		NicknameTh:     req.NicknameTh,
		NicknameEn:     req.NicknameEn,
		Email:          req.Email,
		ContactNumber1: req.ContactNumber1,
		ContactNumber2: req.ContactNumber2,
		LineId:         req.LineId,
		PositionTh:     req.PositionTh,
		PositionEn:     req.PositionEn,
		DepartmentTh:   req.DepartmentTh,
		DepartmentEn:   req.DepartmentEn,
		UpdatedAt:      time.Now(),
	}

	if err := transaction.Where("id = ?", req.Id).Updates(&card).Error; err != nil {
		transaction.Rollback()
		c.logger.ErrorWithID(ctx, log.AppLog, "[Repo:Updated data] update   error ", err)
		return nil, err
	}

	transaction.Commit()

	return card, nil
}

func (c CardRepository) Delete(ctx context.Context, req model.CardRequest) error {
	c.logger.DebugWithID(ctx, log.AppLog, "[Repo:Delete] Called")

	transaction := c.db.Db.Begin()

	card := &entity.Card{}

	r := transaction.Where("id = ?", req.Id).Limit(1).Find(&card)

	if r.RowsAffected == 0 {
		err := errors.New("record not found")
		c.logger.ErrorWithID(ctx, log.AppLog, "[Repo:Delete data] Delete   error ", err)
		return err
	}

	if err := transaction.Where("id = ?", req.Id).Model(&card).Updates(map[string]interface{}{
		"IsActive": false,
	}).Error; err != nil {
		transaction.Rollback()
		c.logger.ErrorWithID(ctx, log.AppLog, "[Repo:Delete data] Delete   error ", err)
		return err
	}

	transaction.Commit()

	return nil
}
