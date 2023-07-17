package controller_v1

import (
	"card-service/internal/config"
	"card-service/internal/model"
	"card-service/internal/service"
	"card-service/internal/util"
	"card-service/internal/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CardController struct {
	config  *config.Configuration
	logger  *log.Logger
	service service.ServiceGateway
}

func (c *CardController) FindAll(ctx *gin.Context) {

	var req model.QueryCard
	if err := ctx.Bind(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:FindAll] Call service Create error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := c.service.CardService.FindAll(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:FindAll] Call service Create error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}
func (c *CardController) Create(ctx *gin.Context) {
	var req model.CreateCardDto
	if err := ctx.Bind(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:Create] Call service Create error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	c.logger.DebugWithID(ctx, log.AppLog, "[CTRL:Create] Called with param : ", util.ConvertStructToJSONString(req))

	res, err := c.service.CardService.Create(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:Create] Call service Create error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}
func (c *CardController) Update(ctx *gin.Context) {

	var req model.UpdateCardDto

	if err := ctx.BindJSON(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:Update] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := ctx.BindUri(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:Update] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	c.logger.DebugWithID(ctx, log.AppLog, "[CTRL:Update] Called with param : ", util.ConvertStructToJSONString(req))

	res, err := c.service.CardService.Update(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:Update] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}
func (c *CardController) Delete(ctx *gin.Context) {

	var req model.CardRequest
	if err := ctx.BindUri(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:Delete] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	c.logger.DebugWithID(ctx, log.AppLog, "[CTRL:Delete] Called with param : ", util.ConvertStructToJSONString(req))

	err := c.service.CardService.Delete(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:Delete] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, true)
	return

}
func (c *CardController) ExportCards(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "")

}
func (c *CardController) ValidateImportCards(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "")

}
func (c *CardController) ImportCards(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "")

}

// NewCardController ...
func NewCardController(sg service.ServiceGateway, l *log.Logger, config *config.Configuration) *CardController {
	return &CardController{
		config:  config,
		logger:  l,
		service: sg,
	}
}
