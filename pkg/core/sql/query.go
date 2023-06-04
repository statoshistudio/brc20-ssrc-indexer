package sql

import (
	"fmt"
	"time"

	"gorm.io/datatypes"

	// _ "github.com/ByteGum/go-ssrc/pkg/core/indexer"
	"gorm.io/gorm"
)

func GetAllAccounts(db *gorm.DB, current int, perPage int) ([]AccountModel, error) {

	if current == 0 {
		current = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	data := []AccountModel{}
	err := db.Limit(perPage).Offset(perPage * (current - 1)).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetAllAccountTokenBalances(db *gorm.DB, current int, perPage int, address string, token string) ([]Brc20TokenAccountBalanceModel, error) {

	if current == 0 {
		current = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	where := ""
	args := []interface{}{}

	if address != "" && token != "" {
		where = fmt.Sprintf(" %s account_address = ? or token_ticker = ?", where)
		args = append(args, address)
		args = append(args, token)
	} else if address != "" {
		where = fmt.Sprintf(" %s account_address = ?", where)
		args = append(args, address)
	} else if token != "" {
		where = fmt.Sprintf(" %s token_ticker = ?", where)
		args = append(args, token)
	}

	data := []Brc20TokenAccountBalanceModel{}
	err := db.Limit(perPage).Offset(perPage*(current-1)).Where(where, args...).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetAllBrc20Tokens(db *gorm.DB, current int, perPage int, address string) ([]Brc20TokenModel, error) {

	if current == 0 {
		current = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	if address != "" {
		address = fmt.Sprintf(" where address = '%s' ", address)
	}
	data := []Brc20TokenModel{}
	err := db.Raw(address).Limit(perPage).Offset(perPage * (current - 1)).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetAllInscriptions(db *gorm.DB, current int, perPage int) ([]InscriptionModel, error) {

	if current == 0 {
		current = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	data := []InscriptionModel{}
	err := db.Limit(perPage).Offset(perPage * (current - 1)).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetAllGenericInscriptions(db *gorm.DB, current int, perPage int) ([]GenericInscriptionModel, error) {

	if current == 0 {
		current = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	data := []GenericInscriptionModel{}
	err := db.Limit(perPage).Offset(perPage * (current - 1)).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetUnitGenericInscription(db *gorm.DB, inscriptionId string) (*GenericInscriptionModel, error) {

	data := GenericInscriptionModel{}
	err := db.First(&data, "inscription_id = ?", inscriptionId).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetUpdatedInscription(db *gorm.DB, inscriptionId string) (*UpdatedInscriptionsModel, error) {

	data := UpdatedInscriptionsModel{}
	err := db.First(&data, "inscription_id = ?", inscriptionId).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func SaveUpdatedInscription(db *gorm.DB, inscriptionId string) (int, error) {

	data := UpdatedInscriptionsModel{InscriptionId: inscriptionId}
	err := db.Create(&data).Error
	if err != nil {
		return -1, err
	}
	return int(data.ID), nil
}

func SaveNewAccount(db *gorm.DB, address string) (int, error) {

	data := AccountModel{Address: address}
	err := db.Create(&data).Error
	if err != nil {
		return -1, err
	}
	return int(data.ID), nil
}

func GetPendingInscriptions(db *gorm.DB) ([]PendingTransferInscriptionModel, error) {

	data := []PendingTransferInscriptionModel{}
	return data, db.Find(&data).Error
}

func GetAllPendingTransactions(db *gorm.DB, current int, perPage int) ([]PendingTransferInscriptionModel, error) {

	if current == 0 {
		current = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	data := []PendingTransferInscriptionModel{}
	err := db.Limit(perPage).Offset(perPage * (current - 1)).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
func GetUpdatedInscriptions(db *gorm.DB, perPage int) ([]UpdatedInscriptionsModel, error) {

	if perPage == 0 {
		perPage = 200
	}

	data := []UpdatedInscriptionsModel{}
	err := db.Limit(perPage).Where("createdAt > ?", datatypes.Date(time.Now().Add(4*time.Minute))).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetConfig(db *gorm.DB, key string) (*ConfigModel, error) {

	data := ConfigModel{}
	err := db.Where(&ConfigModel{Key: key}).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func SetConfig(db *gorm.DB, key string, value string) (*ConfigModel, error) {

	data := ConfigModel{}
	err := db.Where(ConfigModel{Key: key}).Assign(ConfigModel{Value: value}).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
