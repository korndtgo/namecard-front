package error_wrapper

import "errors"

// xxxx
// x[0] = service code oa =1
// x[1] =
// x[2:3] = running number

var (
	ErrorCodeStatus = map[ErrorCode]string{
		"400": "BAD_REQUEST",
		"401": "UNAUTHORIZED",
		"403": "FORBIDEN",
		"404": "NOT FOUND",
		"500": "INTERNAL ERROR",
		"503": "SERVER UNAVAILABLE",

		// RewardItemError
		"5100": "REWARD ITEM ERROR",
		"5101": "INSUFFICIENT STOCK",

		// ClaimOrderError
		"5200": "CLAIM ORDER ERROR ",
		"5201": "LIMITATION EXCEEDED",
		"5202": "INSUFFICIENT_BALANCE",

		// CampaignError
		"5300": "AMPAIGN_ERROR",
		"5301": "CLAIM PHASE_NOT RUNNING",
	}
)

var (
	InternalError = map[ErrorCode]error{
		UNAUTHORIZED: errors.New("unauthorized"),
		NOT_FOUND:    errors.New("not found"),

		REWARD_ITEM_ERROR:    errors.New("reward item error"),
		INSUFFICIENT_STOCK:   errors.New("insufficient stock"),
		CLAIM_ORDER_ERROR:    errors.New("claim order error"),
		LIMITATION_EXCEEDED:  errors.New("limitation exceeded"),
		INSUFFICIENT_BALANCE: errors.New("insufficient balance"),

		// CampaignError
		CAMPAIGN_ERROR:          errors.New("campaign error"),
		CLAIM_PHASE_NOT_RUNNING: errors.New("claim phase not running"),
	}
)

// client error code
const (

	// server error
	BAD_REQUEST        ErrorCode = "400"
	UNAUTHORIZED       ErrorCode = "401"
	FORBIDEN           ErrorCode = "403"
	NOT_FOUND          ErrorCode = "404"
	INTERNAL_ERROR     ErrorCode = "500"
	SERVER_UNAVAILABLE ErrorCode = "503"
)

// RewardItemError
const (
	REWARD_ITEM_ERROR  ErrorCode = "5100"
	INSUFFICIENT_STOCK ErrorCode = "5101"
)

// ClaimOrderError
const (
	CLAIM_ORDER_ERROR    ErrorCode = "5200"
	LIMITATION_EXCEEDED  ErrorCode = "5201"
	INSUFFICIENT_BALANCE ErrorCode = "5202"
)

// CampaignError
const (
	CAMPAIGN_ERROR          ErrorCode = "5300"
	CLAIM_PHASE_NOT_RUNNING ErrorCode = "5301"
)
