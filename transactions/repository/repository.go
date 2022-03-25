package repository

import "database/sql"

type transactionRepository struct {
	TransactionDB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{TransactionDB: db}
}
