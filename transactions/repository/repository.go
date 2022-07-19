package repository

import (
	"payment/models"

	"gorm.io/gorm"
	"gorm.io/hints"
)

type transactionRepository struct {
	TransactionDB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{TransactionDB: db}
}

func (r *transactionRepository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.TransactionDB.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.TransactionDB.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepository) GetByID(ID int64) (models.Transaction, error) {
	var transaction models.Transaction
	// with index
	err := r.TransactionDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&transaction).Error
	// without indexing
	// err := r.TransactionDB.Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetByTrxID(TrxID string) (models.Transaction, error) {
	var transaction models.Transaction
	// with index
	// err := r.TransactionDB.Clauses(hints.UseIndex("idx_status")).Where("trx_id = ?", TrxID).Find(&transaction).Error
	// without indexing
	err := r.TransactionDB.Where("trx_id = ?", TrxID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
