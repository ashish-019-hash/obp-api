package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	instance *sql.DB
	once     sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		instance, err = sql.Open("sqlite3", ":memory:")
		if err != nil {
			log.Fatal("Failed to open database:", err)
		}
		
		if err = instance.Ping(); err != nil {
			log.Fatal("Failed to ping database:", err)
		}
		
		if err = initializeSchema(); err != nil {
			log.Fatal("Failed to initialize schema:", err)
		}
		
		log.Println("SQLite in-memory database initialized successfully")
	})
	return instance
}

func initializeSchema() error {
	return createTables()
}

func createTables() error {
	db := instance
	
	tables := []string{
		`CREATE TABLE IF NOT EXISTS banks (
			bank_id TEXT PRIMARY KEY,
			short_name TEXT NOT NULL,
			full_name TEXT NOT NULL,
			logo TEXT,
			website TEXT,
			bank_routing_scheme TEXT,
			bank_routing_address TEXT
		)`,
		
		`CREATE TABLE IF NOT EXISTS bank_accounts (
			account_id TEXT PRIMARY KEY,
			bank_id TEXT NOT NULL,
			label TEXT NOT NULL,
			number TEXT NOT NULL,
			type TEXT NOT NULL,
			balance_currency TEXT NOT NULL,
			balance_amount TEXT NOT NULL,
			iban TEXT,
			swift_bic TEXT,
			account_routing_scheme TEXT,
			account_routing_address TEXT,
			FOREIGN KEY (bank_id) REFERENCES banks(bank_id)
		)`,
		
		`CREATE TABLE IF NOT EXISTS transactions (
			transaction_id TEXT PRIMARY KEY,
			account_id TEXT NOT NULL,
			counterparty_id TEXT,
			amount_currency TEXT NOT NULL,
			amount_value TEXT NOT NULL,
			description TEXT,
			posted_date TEXT,
			completed_date TEXT,
			new_balance_currency TEXT,
			new_balance_amount TEXT,
			value_date TEXT,
			FOREIGN KEY (account_id) REFERENCES bank_accounts(account_id)
		)`,
		
		`CREATE TABLE IF NOT EXISTS counterparties (
			counterparty_id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			created_by_user_id TEXT,
			this_bank_id TEXT,
			this_account_id TEXT,
			this_view_id TEXT,
			counterparty_id_value TEXT,
			kind TEXT,
			FOREIGN KEY (this_bank_id) REFERENCES banks(bank_id),
			FOREIGN KEY (this_account_id) REFERENCES bank_accounts(account_id)
		)`,
		
		`CREATE TABLE IF NOT EXISTS counterparty_limits (
			limit_id TEXT PRIMARY KEY,
			counterparty_id TEXT NOT NULL,
			max_single_amount_currency TEXT NOT NULL,
			max_single_amount_value TEXT NOT NULL,
			max_monthly_amount_currency TEXT NOT NULL,
			max_monthly_amount_value TEXT NOT NULL,
			FOREIGN KEY (counterparty_id) REFERENCES counterparties(counterparty_id)
		)`,
		
		`CREATE TABLE IF NOT EXISTS users (
			user_id TEXT PRIMARY KEY,
			provider TEXT NOT NULL,
			provider_id TEXT NOT NULL,
			username TEXT,
			email TEXT
		)`,
		
		`CREATE TABLE IF NOT EXISTS customers (
			customer_id TEXT PRIMARY KEY,
			bank_id TEXT NOT NULL,
			customer_number TEXT NOT NULL,
			legal_name TEXT NOT NULL,
			mobile_phone_number TEXT,
			email TEXT,
			face_image_url TEXT,
			date_of_birth TEXT,
			relationship_status TEXT,
			dependents INTEGER,
			dob_of_dependents TEXT,
			highest_education_attained TEXT,
			employment_status TEXT,
			kyc_status BOOLEAN,
			last_ok_date TEXT,
			credit_rating_rating TEXT,
			credit_rating_source TEXT,
			credit_limit_currency TEXT,
			credit_limit_amount TEXT,
			FOREIGN KEY (bank_id) REFERENCES banks(bank_id)
		)`,
		
		`CREATE TABLE IF NOT EXISTS fx_rates (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			bank_id TEXT,
			from_currency TEXT NOT NULL,
			to_currency TEXT NOT NULL,
			rate TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (bank_id) REFERENCES banks(bank_id)
		)`,
		
		`CREATE TABLE IF NOT EXISTS api_metrics (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			consumer_id TEXT,
			user_id TEXT,
			url TEXT NOT NULL,
			duration INTEGER NOT NULL,
			date DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		
		`CREATE TABLE IF NOT EXISTS rate_limits (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			consumer_key TEXT NOT NULL,
			period TEXT NOT NULL,
			period_key TEXT NOT NULL,
			call_count INTEGER DEFAULT 0,
			limit_value INTEGER NOT NULL,
			reset_time DATETIME,
			UNIQUE(consumer_key, period, period_key)
		)`,
	}
	
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
	}
	
	return nil
}
