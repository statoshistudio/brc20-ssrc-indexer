package sql

import (
	"gorm.io/gorm"
)

type AccountModel struct {
	gorm.Model
	ID          int64   `gorm:"primaryKey;autoIncrement:true"`
	Name        *string `json:"name"`
	Address     string  `json:"address"`
	Description *string `json:"description"`
}

type Brc20TokenAccountBalanceModel struct {
	gorm.Model
	ID             int64  `gorm:"primaryKey;autoIncrement:true"`
	TokenTicker    string `json:"token_ticker"`
	AccountAddress string `json:"account_address"`

	AvailableBalance    float64 `json:"available_balance"`
	TransferableBalance float64 `json:"transferable_balance"`
}

type Brc20TokenModel struct {
	gorm.Model
	ID          int64   `gorm:"primaryKey;autoIncrement:true"`
	Ticker      string  `json:"ticker"`
	Address     string  `json:"address"`
	Max         float64 `json:"max"`
	Lim         float64 `json:"lim"`
	Minted      float64 `json:"minted"`
	Description *string `json:"description"`
}

type InscriptionModel struct {
	gorm.Model
	ID            int64  `gorm:"primaryKey;autoIncrement:true"`
	InscriptionId string `gorm:"unique"`
	ErrorMessage  string `json:"error_message"`
	ErrorCode     int    `json:"error_code"`
}

type PendingTransferInscriptionModel struct {
	gorm.Model
	ID             int64  `gorm:"primaryKey;autoIncrement:true"`
	InscriptionId  string `gorm:"unique"`
	GenesisAddress string `json:"genesis_address"`
}

type GenericInscriptionModel struct {
	gorm.Model
	ID                       int64  `gorm:"primaryKey;autoIncrement:true"`
	InscriptionId            string `gorm:"unique"`
	Address                  string `gorm:"address"`
	GenesisAddress           string `gorm:"genesis_address"`
	GenesisFee               int64  `gorm:"genesis_fee"`
	GenesisHeight            int64  `gorm:"genesis_height"`
	InscriptionBody          []byte `gorm:"inscription_body"`
	InscriptionContentLength int64  `gorm:"inscription_content_length"`
	InscriptionContentType   string `gorm:"inscription_content_type"`
	Next                     string `json:"next"`
	Previous                 string `json:"previous"`
	Number                   int64  `json:"number"`
	ScriptPubkey             string `json:"script_pubkey"`
	Value                    int64  `json:"value"`
	OutputAddress            string `json:"output_address"`
	Sat                      string `json:"sat"`
	Satpoint                 string `json:"satpoint"`
	Timestamp                string `json:"timestamp"`
}

type UpdatedInscriptionsModel struct {
	gorm.Model
	ID            int64  `gorm:"primaryKey;autoIncrement:true"`
	InscriptionId string `gorm:"inscription_id"`
}
type ConfigModel struct {
	gorm.Model
	Key   string `gorm:"key;unique"`
	Value string `gorm:"value"`
}
