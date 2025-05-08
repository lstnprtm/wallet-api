package usecase

import (
	"errors"
	"github.com/lstnprtm/wallet-api/internal/domain"
)

type walletUsecase struct {
	repo domain.WalletRepository
}

func NewWalletUsecase(r domain.WalletRepository) domain.WalletUsecase {
	return &walletUsecase{repo: r}
}

func (u *walletUsecase) GetBalance(userID int64) (*domain.Wallet, error) {
	return u.repo.GetWalletByUserID(userID)
}

func (u *walletUsecase) Withdraw(userID int64, amount int64) error {
	wallet, err := u.repo.GetWalletByUserID(userID)
	if err != nil {
		return err
	}
	if wallet.Balance < amount {
		return errors.New("insufficient balance")
	}
	newBalance := wallet.Balance - amount
	if err := u.repo.UpdateBalance(userID, newBalance); err != nil {
		return err
	}
	return u.repo.LogTransaction(userID, amount, "withdraw")
}

func (u *walletUsecase) Deposit(userID int64, amount int64) error {
	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}
	wallet, err := u.repo.GetWalletByUserID(userID)
	if err != nil {
		return err
	}
	newBalance := wallet.Balance + amount
	if err := u.repo.UpdateBalance(userID, newBalance); err != nil {
		return err
	}
	return u.repo.LogTransaction(userID, amount, "deposit")
}

func (u *walletUsecase) GetHistory(userID int64) ([]domain.WalletHistory, error) {
	return u.repo.GetHistory(userID)
}
