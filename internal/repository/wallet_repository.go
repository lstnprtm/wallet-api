package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lstnprtm/wallet-api/internal/domain"
)

type walletRepo struct {
	db *sqlx.DB
}

func NewWalletRepo(db *sqlx.DB) domain.WalletRepository {
	return &walletRepo{db}
}

func (r *walletRepo) GetWalletByUserID(userID int64) (*domain.Wallet, error) {
	var w domain.Wallet
	err := r.db.Get(&w, "SELECT * FROM wallets WHERE user_id = ?", userID)
	return &w, err
}

func (r *walletRepo) UpdateBalance(userID int64, newBalance int64) error {
	_, err := r.db.Exec("UPDATE wallets SET balance = ? WHERE user_id = ?", newBalance, userID)
	return err
}

func (r *walletRepo) GetHistory(userID int64) ([]domain.WalletHistory, error) {
	var h []domain.WalletHistory
	err := r.db.Select(&h, "SELECT * FROM wallet_histories WHERE user_id = ? ORDER BY created_at DESC", userID)
	return h, err
}

func (r *walletRepo) LogTransaction(userID int64, amount int64, txType string) error {
	_, err := r.db.Exec("INSERT INTO wallet_histories (user_id, amount, type) VALUES (?, ?, ?)", userID, amount, txType)
	return err
}
