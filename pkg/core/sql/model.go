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

type GenericInscriptionModel struct {
	gorm.Model
	InscriptionId            string `gorm:"unique"`
	Address                  string `gorm:"address"`
	GenesisAddress           string `gorm:"genesis_address"`
	GenesisFee               int    `gorm:"genesis_fee"`
	GenesisHeight            int    `gorm:"genesis_height"`
	InscriptionBody          []byte `gorm:"inscription_body"`
	InscriptionContentLength int    `gorm:"inscription_content_length"`
	InscriptionContentType   string `gorm:"inscription_content_type"`
	Next                     string `json:"next"`
	Previous                 string `json:"previous"`
	Number                   int    `json:"number"`
	ScriptPubkey             string `json:"script_pubkey"`
	Value                    int    `json:"value"`
	Sat                      string `json:"sat"`
	Satpoint                 string `json:"satpoint"`
	Timestamp                string `json:"timestamp"`
}
