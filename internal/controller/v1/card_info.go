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

type CardInfoController struct {
	config  *config.Configuration
	logger  *log.Logger
	service service.ServiceGateway
}

// NewCardInfoController ...
func NewCardInfoController(sg service.ServiceGateway, l *log.Logger, config *config.Configuration) *CardInfoController {
	return &CardInfoController{
		config:  config,
		logger:  l,
		service: sg,
	}
}

func (c *CardInfoController) GetCard(ctx *gin.Context) {
	var req model.CardRequest
	if err := ctx.BindUri(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:GetCard] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	c.logger.DebugWithID(ctx, log.AppLog, "[CTRL:GetCard] Called with param : ", util.ConvertStructToJSONString(req))

	card, err := c.service.CardInfoService.GetById(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:GetCard] Call service GetCard error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, card)
	return

}
func (c *CardInfoController) GetVCardTH(ctx *gin.Context) {
	var req model.CardRequest
	if err := ctx.BindUri(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:GetVCardTH] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	fileName, err := c.service.CardInfoService.GetByIdTh(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:GetCard] Call service GetCard error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Header("Content-Type", "text/x-vcard")
	ctx.Header("Content-disposition", "attachment; filename=' + `${fileName}.vcf")
	ctx.File(fileName)

	return
}

func (c *CardInfoController) GetVCardEN(ctx *gin.Context) {
	var req model.CardRequest
	if err := ctx.BindUri(&req); err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:GetVCardEN] Call service Update error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	fileName, err := c.service.CardInfoService.GetByIdEn(ctx, req)

	if err != nil {
		c.logger.ErrorWithID(ctx, log.AppLog, "[CTRL:GetVCardEN] Call service GetCard error: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Header("Content-Type", "text/x-vcard")
	ctx.Header("Content-disposition", "attachment; filename=' + `${fileName}.vcf")
	ctx.File(fileName)

	return

}
