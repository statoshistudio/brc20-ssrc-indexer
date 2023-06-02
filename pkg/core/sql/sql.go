package sql

import (
	"fmt"

	"github.com/ByteGum/go-ssrc/utils"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *gorm.DB
var SqlDBErr error

func InitializeDb(file string, migrations []string) (*gorm.DB, error) {
	// db, err := sql.Open("sqlite3", file)
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&ConfigModel{})
	db.AutoMigrate(&InscriptionModel{})
	db.AutoMigrate(&GenericInscriptionModel{})
	db.AutoMigrate(&Brc20TokenModel{})
	db.AutoMigrate(&AccountModel{})
	db.AutoMigrate(&Brc20TokenAccountBalanceModel{})
	db.AutoMigrate(&PendingTransferInscriptionModel{})
	// for _, m := range migrations {
	// 	if _, err := db.Exec(m); err != nil {
	// 		return nil, err
	// 	}
	// }

	return db, err
}

func init() {
	cfg := utils.Config

	SqlDB, SqlDBErr = InitializeDb(fmt.Sprintf("%s/indexer.db", cfg.DataDir), Migrations)
	fmt.Printf("%s/indexer.db = %s\n ", cfg.DataDir, SqlDBErr)

}
