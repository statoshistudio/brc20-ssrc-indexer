package sql

import "gorm.io/gorm"

type AccountModel struct {
	gorm.Model
	Name        *string `json:"name"`
	Address     string  `json:"address"`
	Description *string `json:"description"`
}

type Brc20TokenAccountBalanceModel struct {
	gorm.Model
	TokenTicker    string `json:"token_ticker"`
	AccountAddress string `json:"account_address"`

	AvailableBalance    float64 `json:"available_balance"`
	TransferableBalance float64 `json:"transferable_balance"`
}

type Brc20TokenModel struct {
	gorm.Model
	Ticker      string  `json:"ticker"`
	Address     string  `json:"address"`
	Max         float64 `json:"max"`
	Lim         float64 `json:"lim"`
	Minted      float64 `json:"minted"`
	Description *string `json:"description"`
}

type InscriptionModel struct {
	gorm.Model
	InscriptionId string `gorm:"unique"`
	ErrorMessage  string `json:"error_message"`
	ErrorCode     int    `json:"error_code"`
}

type PendingTransferInscriptionModel struct {
	gorm.Model
	InscriptionId  string `gorm:"unique"`
	GenesisAddress string `json:"genesis_address"`
}
