package sql

import (
	"github.com/ByteGum/go-ssrc/utils"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *gorm.DB
var SqlDBErr error
var cfg utils.Configuration

func InitializeDb(driver string, dsn string, migrations []string) (*gorm.DB, error) {
	// db, err := sql.Open("sqlite3", file)
	var dialector gorm.Dialector
	switch driver {
	case "mysql":
		dialector = mysql.Open(dsn)
	default:
		dialector = sqlite.Open(dsn)
	}
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel()),
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
	db.AutoMigrate(&UpdatedInscriptionsModel{})
	// for _, m := range migrations {
	// 	if _, err := db.Exec(m); err != nil {
	// 		return nil, err
	// 	}
	// }

	return db, err
}

func init() {
	cfg := utils.Config

	SqlDB, SqlDBErr = InitializeDb(cfg.DbDriver, cfg.DbDSN, Migrations)
	if SqlDBErr != nil {
		panic(SqlDBErr)
	}
}

func logLevel() logger.LogLevel {
	if cfg.LogLevel == "info" {
		return logger.Info
	}
	return logger.Silent
}
