package constant_v1

const (
	WAITING_FOR_INVEST = "WAITING_FOR_INVEST"
	BUY                = "BUY"
	ORDERED            = "ORDERED"
	FAILED             = "FAILED"
	UNSUCCESS          = "UNSUCCESS"
	SUCCESS            = "SUCCESS"
	ALLOTTED           = "ALLOTTED"
	CANCELLED          = "CANCELLED"
	COMPLETED          = "COMPLETED"
	REFUND             = "REFUND"
	REFUNDED           = "REFUNDED"
	DEPOSIT            = "DEPOSIT"
	WITHDRAW           = "WITHDRAW"
	TRANSFERRED        = "TRANSFERRED"
	CLOSED             = "CLOSED"
	//digitSaOrderType
	PRSF = "PRSF" //First Investmnet
	PRSR = "PRSR" //Next Investment
	PRSD = "PRSD" //DCA Investment
	PRRR = "PRRR" //SELL Investment
	PRSI = "PRSI" //Reinvest Investment
	PRSU = "PRSU" //Retransaction Investment
	PRRA = "PRRA" //Robo-RED-Auto
	PRRM = "PRRM" //Robo-RED-Manual
	PRSA = "PRSA" //Robo-SUB-Auto
	PRSM = "PRSM" //Robo-SUB-Manual

	PREPARING       = "PREPARING"
	READY           = "READY"
	ROBO_CASH_READY = "ROBO_CASH_READY"
	SELL            = "SELL"
	DCA             = "DCA"

	CASH = "CASH"
	UNIT = "UNIT"
	MIX  = "MIX"
	AMT  = "AMT"

	PRE_ACTIVE = "PRE_ACTIVE"
	ACTIVE     = "ACTIVE"
	NEW        = "NEW"

	PRACTICAL = "PRACTICAL"

	DCA_STATUS_ACTIVE    = "ACTIVE"
	DCA_STATUS_INACTIVE  = "INACTIVE"
	DCA_STATUS_CANCELLED = "CANCELLED"

	SCHEDULE_TYPE_W       = "W"
	SCHEDULE_TYPE_M       = "M"
	SCHEDULE_TYPE_ALL     = "ALL"
	DCA_STATUS_WAITING_TH = "รอดำเนินการ"
	DCA_STATUS_WAITING_EN = "waiting to be processed"

	ALL  = "ALL"
	EACH = "EACH"
	SUB  = "SUB"
	RED  = "RED"

	NOTI_NAME = "robo-conductor-service"
	NOTI_FROM = "robo-conductor"

	DURATION_ONE_MONTH   = "ONE_MONTH"
	DURATION_THREE_MONTH = "THREE_MONTH"
	DURATION_SIX_MONTH   = "SIX_MONTH"
	DURATION_ONE_YEAR    = "ONE_YEAR"
	DURATION_THREE_YEAR  = "THREE_YEAR"
	DURATION_FIVE_YEAR   = "FIVE_YEAR"
	DURATION_MAX         = "MAX"
	REBALANCE            = "REBALANCE"
	REBALANCE_CASH       = "REBALANCE_CASH"
	REBALANCING          = "REBALANCING"

	REINVEST      = "REINVEST"
	RETRANSACTION = "RETRANSACTION"

	LAYOUT_ISO     = "20060102"
	CURRENT        = "CURRENT"
	RECOMMENDATION = "RECOMMENDATION"

	DIVIDEND       = "DIVIDEND"
	UNDEFINED_CASH = "UNDEFINED-CASH"
	INTEREST       = "INTEREST"
	AR             = "AR"

	DESC = "DESC"
	ASC  = "ASC"

	ACCOUNT_TYPE_S = "S"
)
