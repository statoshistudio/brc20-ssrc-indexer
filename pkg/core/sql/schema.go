package sql

const File string = "./data/sql/indexer.db"

const CreateAccountTable string = `
CREATE TABLE IF NOT EXISTS accounts (
	id INTEGER NOT NULL PRIMARY KEY,
	name VARCHAR(255),
    address VARCHAR(255) UNIQUE,
	description TEXT
	);
`

const CreateBrc20AccountBalanceTable string = `
CREATE TABLE IF NOT EXISTS brc_20_account_balances (
	id INTEGER NOT NULL PRIMARY KEY,
	token_ticker VARCHAR(255) NOT NULL,
    account_address VARCHAR(255) NOT NULL,
	balance DOUBLE DEFAULT 0,
	UNIQUE(token_ticker, account_address) ON CONFLICT REPLACE
	);
`

const CreateBrc20TokenTable string = `
CREATE TABLE IF NOT EXISTS brc_20_tokens (
	id INTEGER NOT NULL PRIMARY KEY,
	ticker VARCHAR(255) UNIQUE,
    address VARCHAR(255),
    max INTEGER,
    lim INTEGER,
    minted INTEGER DEFAULT 0,
	description TEXT
	);
`

const CreateInscriptionTable string = `
CREATE TABLE IF NOT EXISTS inscriptions (
	id INTEGER NOT NULL PRIMARY KEY,
	address VARCHAR(255),
	genesis_fee INTEGER,
	genesis_height INTEGER,
	inscription_id VARCHAR(255) UNIQUE,
	next VARCHAR(255),
	number VARCHAR(255),
	script_pubkey VARCHAR(255),
	value VARCHAR(255),
	previous VARCHAR(255),
	sat VARCHAR(255),
	satpoint VARCHAR(255),
	timestamp DATETIME,
	content Text,
	content_length INTEGER,
	content_type VARCHAR(255)
	);
`

var Migrations = []string{CreateAccountTable, CreateBrc20AccountBalanceTable, CreateBrc20TokenTable, CreateInscriptionTable}
