package tests

import (
	"time"
)

// ================= Begin Port(port_service) =================
type Port struct {
	AccountID   string    `gorm:"primarykey;column:account_id;type:varchar(36);"`
	CID         string    `gorm:"index;column:cid;type:varchar(36);"`
	AccountType string    `gorm:"type:varchar(16);"`
	Status      string    `gorm:"type:varchar(36);"`
	Name        string    `gorm:"type:varchar(128);"`
	PortClass   *string   `gorm:"index;type:varchar(36);"`
	RefNo       *string   `gorm:"index;type:varchar(128);"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// TableName Rename
func (Port) TableName() string {
	return "tdc_port"
}

// ================= End Port(port_service) =================

// ================= Begin User(port_service) =================
type User struct {
	UserId                    string    `gorm:"uniqueIndex;type:varchar(64)" json:"userId"`
	CID                       string    `gorm:"primarykey;type:varchar(36);column:cid" json:"cid"`
	SettlementAccountID       string    `gorm:"type:varchar(64)" json:"settlementAccountId"`
	EnFirstName               string    `gorm:"type:nvarchar(256)" json:"enFirstName"`
	EnLastName                string    `gorm:"type:nvarchar(256)" json:"enLastName"`
	ThFirstName               string    `gorm:"type:nvarchar(256)" json:"thFirstName"`
	ThLastName                string    `gorm:"type:nvarchar(256)" json:"thLastName"`
	IdentificationCardType    string    `gorm:"type:nvarchar(256)" json:"identificationCardType"`
	CardNumber                string    `gorm:"type:nvarchar(256)" json:"cardNumber"`
	CardExpiryDate            string    `gorm:"type:nvarchar(256)" json:"cardExpiryDate"`
	BirthDate                 string    `gorm:"type:nvarchar(256)" json:"birthDate"`
	SuitabilityRiskLevel      string    `gorm:"type:nvarchar(256)" json:"suitabilityRiskLevel"`
	SuitabilityEvaluationDate string    `gorm:"type:nvarchar(256)" json:"suitabilityEvaluationDate"`
	MaxBuyPerDay              float64   `gorm:"type:decimal(18,2)" json:"maxBuyPerDay"`
	VulnerableFlag            *string   `gorm:"type:varchar(1)" json:"vulnerableFlag"`
	VulnerableDetail          *string   `gorm:"type:varchar(255)" json:"vulnerableDetail"`
	CreatedAt                 time.Time `gorm:"autoCreateTime"`
	UpdatedAt                 time.Time `gorm:"autoUpdateTime"`
}

// TableName Rename
func (User) TableName() string {
	return "tdc_user"
}

// ================= End User(port_service) =================

// ================= Begin OrderTxns =================
type OrderTxns struct {
	OrderID        uint64    `gorm:"primarykey; autoIncrement"`
	AccountID      string    `gorm:"type:varchar(32)"`
	OrderType      string    `gorm:"type:varchar(32)"`
	Status         string    `gorm:"type:varchar(16)"`
	ExposureType   string    `gorm:"column:exposure_type; type:varchar(16) NOT NULL"`
	Amount         *float64  `gorm:"type:decimal(18,2)"`
	Unit           *float64  `gorm:"type:decimal(18,4)"`
	ErrorMessage   *string   `gorm:"type:varchar(128)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	DcaDiyId       *string   `gorm:"type:varchar(255)"`
	ReferanceId    string    `gorm:"type:varchar(128);column:referance_id;"`
	CallbackUrl    string    `gorm:"type:nvarchar(512);column:callback_url;"`
	IsLocked       *bool     `gorm:"column:is_locked;type:bit;default:0"`
	Retryable      *bool     `gorm:"column:retryable;type:bit;default:1"`
	CancelDateTime string    `gorm:"column:cancel_date_time;type:VARCHAR(16)"`
}

// TableName Rename
func (OrderTxns) TableName() string {
	return "tdc_order_txns"
}

// ================= End OrderTxns =================

// ================= Begin PortTxns =================
type PortTxns struct {
	PortTxnId                          uint64 `gorm:"primarykey;autoIncrement"`
	OrderId                            uint64
	SaOrderReferenceNo                 string    `gorm:"uniqueIndex; type:varchar(36)"`
	AccountId                          string    `gorm:"type:varchar(36)"`
	UnitholderId                       string    `gorm:"type:varchar(24)"`
	Status                             string    `gorm:"type:varchar(12)"`
	SaCode                             string    `gorm:"type:varchar(24)"`
	AmcCode                            string    `gorm:"type:varchar(24)"`
	TransactionCode                    string    `gorm:"type:varchar(8)"`
	FundCode                           string    `gorm:"type:varchar(36)"`
	OverrideRiskProfile                string    `gorm:"type:varchar(1)"`
	OverrideFxRisk                     string    `gorm:"type:varchar(1)"`
	RedemptionType                     string    `gorm:"type:varchar(8)"`
	Amount                             *float64  `gorm:"type:decimal(18,2)"`
	Unit                               *float64  `gorm:"type:decimal(18,4)"`
	CounterFundCode                    string    `gorm:"type:varchar(36)"`
	PaymentType                        *string   `gorm:"type:varchar(12)"`
	BankCode                           *string   `gorm:"type:varchar(8)"`
	BankAccount                        *string   `gorm:"type:varchar(24)"`
	SellAllUnitFlag                    *string   `gorm:"type:varchar(1)"`
	SettlementBankCode                 *string   `gorm:"type:varchar(8)"`
	SettlementBankAccount              *string   `gorm:"type:varchar(24)"`
	IcLicense                          string    `gorm:"type:varchar(16)"`
	Channel                            string    `gorm:"type:varchar(8)"`
	ForceEntry                         string    `gorm:"type:varchar(1)"`
	EffectiveDate                      string    `gorm:"type:varchar(12)"`
	TransactionDateTime                string    `gorm:"type:varchar(24)"`
	TransactionId                      string    `gorm:"index;type:varchar(24)"`
	CreditCardNo                       *string   `gorm:"type:varchar(24)"`
	CreditCardIssuer                   *string   `gorm:"type:varchar(36)"`
	ChequeNo                           *string   `gorm:"type:varchar(16)"`
	ChequeDate                         *string   `gorm:"type:varchar(12)"`
	LtfCondition                       *string   `gorm:"type:varchar(1)"`
	ReasonToSellLtfRmf                 *string   `gorm:"type:varchar(1)"`
	RmfCapitalGainWithholdingTaxChoice *string   `gorm:"type:varchar(1)"`
	RmfCapitalAmountRedeemChoice       *string   `gorm:"type:varchar(1)"`
	AutoRedeemFundCode                 *string   `gorm:"type:varchar(36)"`
	AmcOrderReferenceNo                *string   `gorm:"type:varchar(36)"`
	AllotmentDate                      *string   `gorm:"type:varchar(12)"`
	AllottedNav                        *float64  `gorm:"type:decimal(18,4)"`
	AllottedAmount                     *float64  `gorm:"type:decimal(18,2)"`
	AllottedUnit                       *float64  `gorm:"type:decimal(18,4)"`
	Fee                                *float64  `gorm:"type:decimal(18,2)"`
	WithholdingTax                     *float64  `gorm:"type:decimal(18,2)"`
	Vat                                *float64  `gorm:"type:decimal(18,2)"`
	BrokerageFee                       *float64  `gorm:"type:decimal(18,2)"`
	WithholdingTaxForLtfRmf            *float64  `gorm:"type:decimal(18,2)"`
	AmcPayDate                         *string   `gorm:"type:varchar(12)"`
	RegistrarTransactionFlag           *string   `gorm:"type:varchar(1)"`
	RejectReason                       *string   `gorm:"type:varchar(64)"`
	ChqBranch                          *string   `gorm:"type:varchar(12)"`
	TaxInvoiceNo                       *string   `gorm:"type:varchar(64)"`
	AmcRelatedOrderReferenceNo         *string   `gorm:"type:varchar(36)"`
	IcCode                             *string   `gorm:"type:varchar(16)"`
	BrokerageFeeVat                    *float64  `gorm:"type:decimal(18,2)"`
	ApprovalCode                       *string   `gorm:"type:varchar(36)"`
	NavDate                            *string   `gorm:"type:varchar(12)"`
	CollateralAccount                  *string   `gorm:"type:varchar(36)"`
	BranchNo                           *string   `gorm:"type:varchar(8)"`
	ErrorMessage                       *string   `gorm:"type:varchar(128)"`
	ManualCancel                       string    `gorm:"type:varchar(1);default:N"`
	CallbackUrl                        string    `gorm:"type:nvarchar(512);column:callback_url;"`
	CreatedAt                          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName Rename
func (PortTxns) TableName() string {
	return "tdc_port_txns"
}

// ================= End PortTxns =================
