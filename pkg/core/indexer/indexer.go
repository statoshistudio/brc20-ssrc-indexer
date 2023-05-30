package indexer

import (
	// "errors"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	// "github.com/ByteGum/go-ssrc/pkg/core/sql"
	"github.com/ByteGum/go-ssrc/pkg/core/sql"
	utils "github.com/ByteGum/go-ssrc/utils"
	"gorm.io/gorm"
)

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type IndexerService struct {
	Ctx                    *context.Context
	Cfg                    *utils.Configuration
	ClientHandshakeChannel *chan *utils.ClientHandshake
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type InscriptionResponses struct {
	Data []string     `json:"data"`
	Meta MetaResponse `json:"meta"`
}

type MetaResponse struct {
	Status     bool               `json:"success"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	Current *int64 `json:"current"`
	Next    *int64 `json:"next"`
	Prev    *int64 `json:"prev"`
}

type InscriptionResponse struct {
	Address           string            `json:"address"`
	GenesisAddress    string            `json:"genesis_address"`
	GenesisFee        *int64            `json:"genesis_fee"`
	GenesisHeight     *int64            `json:"genesis_height"`
	InscriptionId     string            `json:"inscription_id"`
	Next              *string           `json:"next"`
	Number            *int64            `json:"number"`
	Previous          string            `json:"previous"`
	Sat               string            `json:"sat"`
	Satpoint          string            `json:"satpoint"`
	Timestamp         string            `json:"timestamp"`
	InscriptionOutput InscriptionOutput `json:"output"`
	Inscription       Inscription       `json:"inscription"`
}

type InscriptionOutput struct {
	Address      string `json:"address"`
	ScriptPubkey string `json:"script_pubkey"`
	Value        int64  `json:"value"`
}

type TxInput struct {
	PreviousOutput string `json:"previous_output"`
}

type Inscription struct {
	Body          []byte `json:"body"`
	ContentLength *int64 `json:"content_length"`
	ContentType   string `json:"content_type"`
}
type StaticInscriptionStructure struct {
	P string `json:"p"`
}
type InscriptionStructure struct {
	P      string `json:"p"`
	Op     string `json:"op"`
	Ticker string `json:"tick"`
	Max    string `json:"max"`
	Lim    string `json:"lim"`
	Amt    string `json:"amt"`
}

type Transaction struct {
	Data TransactionData `json:"data"`
}

type TransactionData struct {
	Blockhash   *string          `json:"blockhash"`
	Inscription string           `json:"inscription"`
	Number      int              `json:"number"`
	Transaction TransactionDatum `json:"transaction"`
}

type TransactionDatum struct {
	Blockhash   *string             `json:"blockhash"`
	Inscription string              `json:"inscription"`
	LockTime    int                 `json:"lock_time"`
	Version     int                 `json:"version"`
	Input       []TxInput           `json:"input"`
	Output      []InscriptionOutput `json:"output"`
}

func (i *Inscription) GetContent() InscriptionStructure {
	var inscriptionStructure InscriptionStructure
	err := json.Unmarshal(i.Body, &inscriptionStructure)
	if err != nil {
		log.Println("Errr err:", err)
	}
	return inscriptionStructure
}

// func (g *sql.GenericInscriptionModel) GetContent() InscriptionStructure {
// 	var inscriptionStructure InscriptionStructure
// 	err := json.Unmarshal([]byte(g.InscriptionBody), &inscriptionStructure)
// 	if err != nil {
// 		log.Println("Errr err:", err)
// 	}
// 	return inscriptionStructure
// }

func GetDataFromServer(mainCtx *context.Context, page *int) (*InscriptionResponses, error) {
	// utils.Logger.Infof("GetDataFromServer page  : %d,  %d", *page, *index)

	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	utils.Logger.Infof("GetDataFromServer page neV : %d,  %s", *page, fmt.Sprintf("%s%s%d", cfg.OrdinalApi, "/api/inscriptions/", *page))
	resp, err1 := http.Get(fmt.Sprintf("%s%s%d", cfg.OrdinalApi, "/api/inscriptions/", *page))
	if err1 != nil {
		return nil, err1
	}
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	var inscriptionResponses InscriptionResponses

	err3 := json.Unmarshal(body, &inscriptionResponses)
	// rr, _ := json.Marshal(inscriptionResponses.Data)
	// utils.Logger.Infof("@@@@inscriptionResponses : %s", rr)
	if err3 != nil {
		return nil, err3
	}

	return &inscriptionResponses, nil
}

func GetUnitDataByIdFromServer(mainCtx *context.Context, id string) (*InscriptionResponse, error) {
	// utils.Logger.Infof("+++++++ GetUnitDataByIdFromServer id : %+s", id)
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	resp, err1 := http.Get(fmt.Sprintf("%s%s%s", cfg.OrdinalApi, "/api/inscription/", id))
	if err1 != nil {
		return nil, err1
	}
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	var inscriptionResponse InscriptionResponse
	err3 := json.Unmarshal(body, &inscriptionResponse)
	if err3 != nil {
		return nil, err3
	}
	// utils.Logger.Infof("<><><><><><>GetUnitDataByIdFromServer Number:  %d, = %s : %+s ", *inscriptionResponse.Number, id, string(inscriptionResponse.Inscription.ContentType))
	return &inscriptionResponse, nil
}

func GetUnitTransactionByIdFromServer(mainCtx *context.Context, id string) (*Transaction, error) {
	// utils.Logger.Infof("+++++++ GetUnitDataByIdFromServer id : %+s", id)
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	resp, err1 := http.Get(fmt.Sprintf("%s%s%s", cfg.OrdinalApi, "/api/tx/", id))
	if err1 != nil {
		return nil, err1
	}
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	var transaction Transaction
	err3 := json.Unmarshal(body, &transaction)
	if err3 != nil {
		return nil, err3
	}
	utils.Logger.Infof("------GetUnitTransactionByIdFromServer Type = %s : %+s ", id, string(transaction.Data.Transaction.Output[0].Address))
	return &transaction, nil
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func SaveUnitInscription(db *gorm.DB, inscriptionResponse *InscriptionResponse) error {

	data := sql.InscriptionModel{
		InscriptionId: inscriptionResponse.InscriptionId,
	}
	utils.Logger.Infof("@@@@SaveUnitInscription Type = %s  ", string(inscriptionResponse.Inscription.ContentType))
	return db.Create(&data).Error
}

func SaveGenericInscription(db *gorm.DB, inscriptionResponse *InscriptionResponse) error {

	next := ""
	if inscriptionResponse.Next != nil {
		next = *inscriptionResponse.Next
	}
	data := sql.GenericInscriptionModel{
		InscriptionId:            inscriptionResponse.InscriptionId,
		Address:                  inscriptionResponse.Address,
		GenesisAddress:           inscriptionResponse.GenesisAddress,
		GenesisFee:               int(*inscriptionResponse.GenesisFee),
		GenesisHeight:            int(*inscriptionResponse.GenesisHeight),
		InscriptionBody:          inscriptionResponse.Inscription.Body,
		InscriptionContentLength: int(*inscriptionResponse.Inscription.ContentLength),
		InscriptionContentType:   inscriptionResponse.Inscription.ContentType,
		Next:                     next,
		Previous:                 inscriptionResponse.Previous,
		Number:                   int(*inscriptionResponse.Number),
		ScriptPubkey:             inscriptionResponse.InscriptionOutput.ScriptPubkey,
		Value:                    int(inscriptionResponse.InscriptionOutput.Value),
		Sat:                      inscriptionResponse.Sat,
		Satpoint:                 inscriptionResponse.Satpoint,
		Timestamp:                inscriptionResponse.Timestamp,
	}
	utils.Logger.Infof("@@@@SaveGenericInscription Type = %s  ", string(inscriptionResponse.Inscription.ContentType))
	return db.Create(&data).Error
}

func SaveNewToken(db *gorm.DB, inscriptionStructure *InscriptionStructure, address string) error {

	utils.Logger.Infof("@@@@SaveNewToken OP = %s : %s  ", string(inscriptionStructure.Op), string(inscriptionStructure.Ticker))

	max, _ := strconv.ParseFloat(inscriptionStructure.Max, 64)
	lim, _ := strconv.ParseFloat(inscriptionStructure.Lim, 64)
	data := sql.Brc20TokenModel{Address: address, Ticker: inscriptionStructure.Ticker, Max: max, Lim: lim}

	return db.Create(&data).Error
}

func PerformMintOperation(db *gorm.DB, inscriptionStructure *InscriptionStructure, inscription InscriptionResponse) error {

	err := db.Transaction(func(tx *gorm.DB) error {

		var brc20TokenModel sql.Brc20TokenModel
		err := tx.First(&brc20TokenModel, "ticker=?", inscriptionStructure.Ticker).Error

		if err != nil {
			tx.Model(&sql.InscriptionModel{}).Where("inscription_id = ?", inscription.InscriptionId).Updates(map[string]interface{}{"ErrorCode": utils.EC4004.Code, "ErrorMessage": utils.EC4004.Message})
			return nil
		}

		ogAmt, _ := strconv.ParseFloat(inscriptionStructure.Amt, 64)
		if ogAmt > brc20TokenModel.Lim {

			tx.Model(&sql.InscriptionModel{}).Where("inscription_id = ?", inscription.InscriptionId).Updates(map[string]interface{}{"ErrorCode": utils.EC4005.Code, "ErrorMessage": utils.EC4005.Message})

			return nil
		}

		if brc20TokenModel.Minted >= brc20TokenModel.Max {

			tx.Model(&sql.InscriptionModel{}).Where("inscription_id = ?", inscription.InscriptionId).Updates(map[string]interface{}{"ErrorCode": utils.EC3001.Code, "ErrorMessage": utils.EC3001.Message})

			return nil
		}

		var brc20TokenAccountBalanceModel sql.Brc20TokenAccountBalanceModel
		err = tx.FirstOrCreate(&brc20TokenAccountBalanceModel, sql.Brc20TokenAccountBalanceModel{TokenTicker: inscriptionStructure.Ticker, AccountAddress: inscription.GenesisAddress}).Error
		if err != nil {
			// return any error will rollback
			utils.Logger.Infof("@@@@brc20TokenAccountBalanceModel err = %s  ", err)
			return err
		}

		// utils.Logger.Infof("@@@@ogAmt-(brc20TokenModel.Max-brc20TokenModel.Minted) = %f - %f   ", ogAmt, (brc20TokenModel.Max - brc20TokenModel.Minted))
		if ogAmt-(brc20TokenModel.Max-brc20TokenModel.Minted) > 0 {
			ogAmt = brc20TokenModel.Max - brc20TokenModel.Minted
		}
		err = tx.Model(&brc20TokenAccountBalanceModel).Update("available_balance", gorm.Expr("available_balance + ?", ogAmt)).Error
		// err = tx.Model(&brc20TokenAccountBalanceModel).Updates(map[string]interface{}{"available_balance": availableBalance, "token_ticker": inscriptionStructure.Ticker, "account_address": inscription.Address}).Error

		if err != nil {
			// return any error will rollback
			return err
		}

		err = tx.Model(&brc20TokenModel).Update("minted", gorm.Expr("minted + ?", ogAmt)).Error

		if err != nil {
			// return any error will rollback
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return err
}

func PerformTransferOperation(db *gorm.DB, inscriptionStructure *InscriptionStructure, inscription InscriptionResponse, genesisAddress string) error {

	err := db.Transaction(func(tx *gorm.DB) error {

		var brc20TokenModel sql.Brc20TokenModel
		err := tx.First(&brc20TokenModel, "ticker=?", inscriptionStructure.Ticker).Error

		if err != nil {
			tx.Model(&sql.InscriptionModel{}).Where("inscription_id = ?", inscription.InscriptionId).Updates(map[string]interface{}{"ErrorCode": utils.EC4004.Code, "ErrorMessage": utils.EC4004.Message})
			return nil
		}

		ogAmt, _ := strconv.ParseFloat(inscriptionStructure.Amt, 64)

		var brc20TokenAccountBalanceModel sql.Brc20TokenAccountBalanceModel
		err = tx.FirstOrCreate(&brc20TokenAccountBalanceModel, " token_ticker=? and account_address=? ", inscriptionStructure.Ticker,
			inscription.GenesisAddress).Error
		if err != nil {
			// return any error will rollback

			return err
		}

		// availableBalance := brc20TokenAccountBalanceModel.OverallBalance - brc20TokenAccountBalanceModel.TransferableBalance
		// utils.Logger.Infof("@@@@brc20TokenAccountBalanceModel token_ticker=%s and account_address=%s ", inscriptionStructure.Ticker, inscription.GenesisAddress)
		// utils.Logger.Infof("@@@@ogAmt > availableBalance %f = %f   brc20TokenAccountBalanceModel.AvailableBalance - brc20TokenAccountBalanceModel.TransferableBalance  %f ", ogAmt, brc20TokenAccountBalanceModel.AvailableBalance, brc20TokenAccountBalanceModel.TransferableBalance)
		// utils.Logger.Infof("@@@@Pass 0.0 >>> %s ", inscriptionStructure.Ticker)
		if ogAmt > brc20TokenAccountBalanceModel.AvailableBalance {
			tx.Model(&sql.InscriptionModel{}).Where("inscription_id = ?", inscription.InscriptionId).Updates(map[string]interface{}{"ErrorCode": utils.EC4001.Code, "ErrorMessage": utils.EC4001.Message})
			return nil

		}
		// utils.Logger.Infof("@@@@Pass %s 0 >>>", inscriptionStructure.Ticker)
		err = tx.Model(&brc20TokenAccountBalanceModel).Updates(map[string]interface{}{"transferable_balance": gorm.Expr("transferable_balance + ?", ogAmt), "available_balance": gorm.Expr("available_balance - ?", ogAmt), "token_ticker": inscriptionStructure.Ticker, "account_address": inscription.GenesisAddress}).Error
		// utils.Logger.Infof("@@@@Pass %s 1 >>>", inscriptionStructure.Ticker)
		if err != nil {
			// return any error will rollback
			// utils.Logger.Infof("@@@@after  %s >>>  %s ", inscriptionStructure.Ticker, err)
			return err
		}
		// utils.Logger.Infof("@@@@Pass  %s 2 >>>", inscriptionStructure.Ticker)

		pendingTransferInscriptionModel := sql.PendingTransferInscriptionModel{
			InscriptionId:  inscription.InscriptionId,
			GenesisAddress: genesisAddress,
		}
		err = tx.Create(&pendingTransferInscriptionModel).Error

		if err != nil {
			// return any error will rollback
			utils.Logger.Infof("@@@@pendingTransferInscriptionModel err = %s  ", err)
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return err
}

func CreditPendingOperation(db *gorm.DB, inscriptionStructure *InscriptionStructure, inscription sql.GenericInscriptionModel, recieverAddress string) error {

	err := db.Transaction(func(tx *gorm.DB) error {

		var brc20TokenModel sql.Brc20TokenModel
		err := tx.First(&brc20TokenModel, "ticker=?", inscriptionStructure.Ticker).Error

		if err != nil {
			tx.Model(&sql.InscriptionModel{}).Where("inscription_id = ?", inscription.InscriptionId).Updates(map[string]interface{}{"ErrorCode": utils.EC4004.Code, "ErrorMessage": utils.EC4004.Message})
			return nil
		}

		ogAmt, _ := strconv.ParseFloat(inscriptionStructure.Amt, 64)

		var genesisBrc20TokenAccountBalanceModel sql.Brc20TokenAccountBalanceModel
		var recieverBrc20TokenAccountBalanceModel sql.Brc20TokenAccountBalanceModel

		err = tx.First(&genesisBrc20TokenAccountBalanceModel, " token_ticker=? and account_address=? ", inscriptionStructure.Ticker,
			inscription.GenesisAddress).Error
		if err != nil {
			// return any error will rollback
			return err
		}
		err = tx.FirstOrCreate(&recieverBrc20TokenAccountBalanceModel, " token_ticker=? and account_address=? ", inscriptionStructure.Ticker,
			recieverAddress).Error
		if err != nil {
			// return any error will rollback
			return err
		}

		utils.Logger.Infof("@@@@Pass %s 0 >>>", inscriptionStructure.Ticker)
		err = tx.Model(&genesisBrc20TokenAccountBalanceModel).Updates(map[string]interface{}{"transferable_balance": gorm.Expr("transferable_balance - ?", ogAmt), "token_ticker": inscriptionStructure.Ticker, "account_address": inscription.GenesisAddress}).Error
		utils.Logger.Infof("@@@@Pass %s 1 >>>", inscriptionStructure.Ticker)
		if err != nil {
			// return any error will rollback
			utils.Logger.Infof("@@@@after  %s >>>  %s ", inscriptionStructure.Ticker, err)
			return err
		}
		err = tx.Model(&recieverBrc20TokenAccountBalanceModel).Updates(map[string]interface{}{"available_balance": gorm.Expr("available_balance + ?", ogAmt), "token_ticker": inscriptionStructure.Ticker, "account_address": inscription.GenesisAddress}).Error
		utils.Logger.Infof("@@@@Pass %s 1.1 >>>", inscriptionStructure.Ticker)
		if err != nil {
			// return any error will rollback
			utils.Logger.Infof("@@@@after  %s >>>  %s ", inscriptionStructure.Ticker, err)
			return err
		}
		utils.Logger.Infof("@@@@Pass  %s 2 >>>", inscriptionStructure.Ticker)

		// InscriptionId:  inscription.InscriptionId,
		// 	GenesisAddress: inscription.GenesisAddress,
		err = tx.Delete(&sql.PendingTransferInscriptionModel{}, " inscription_id=? and genesis_address=? ", inscription.InscriptionId, inscription.GenesisAddress).Error
		if err != nil {
			// return any error will rollback
			utils.Logger.Infof("@@@@pendingTransferInscriptionModel err = %s  ", err)
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return err
}

func HandleCallback(db *gorm.DB, inscriptionResponse InscriptionResponse) (*sql.GenericInscriptionModel, error) {

	data := sql.GenericInscriptionModel{}
	err := db.First(&data, "inscription_id = ?", inscriptionResponse.InscriptionId).Error
	if err != nil {
		return nil, err
	}

	next := ""
	if inscriptionResponse.Next != nil {
		next = *inscriptionResponse.Next
	}
	err = db.Model(&data).Updates(sql.GenericInscriptionModel{
		Address:                  inscriptionResponse.Address,
		GenesisAddress:           inscriptionResponse.GenesisAddress,
		GenesisFee:               int(*inscriptionResponse.GenesisFee),
		GenesisHeight:            int(*inscriptionResponse.GenesisHeight),
		InscriptionBody:          inscriptionResponse.Inscription.Body,
		InscriptionContentLength: int(*inscriptionResponse.Inscription.ContentLength),
		InscriptionContentType:   inscriptionResponse.Inscription.ContentType,
		Next:                     next,
		Previous:                 inscriptionResponse.Previous,
		Number:                   int(*inscriptionResponse.Number),
		ScriptPubkey:             inscriptionResponse.InscriptionOutput.ScriptPubkey,
		Value:                    int(inscriptionResponse.InscriptionOutput.Value),
		Sat:                      inscriptionResponse.Sat,
		Satpoint:                 inscriptionResponse.Satpoint,
		Timestamp:                inscriptionResponse.Timestamp,
	}).Error

	return &data, nil
}
