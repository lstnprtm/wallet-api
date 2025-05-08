package domain

type Wallet struct {
	ID      int64 `db:"id"`
	UserID  int64 `db:"user_id"`
	Balance int64 `db:"balance"`
}

type WalletHistory struct {
	ID        int64  `db:"id"`
	UserID    int64  `db:"user_id"`
	Amount    int64  `db:"amount"`
	Type      string `db:"type"`
	CreatedAt string `db:"created_at"`
}

type WalletRepository interface {
	GetWalletByUserID(userID int64) (*Wallet, error)
	UpdateBalance(userID int64, newBalance int64) error
	GetHistory(userID int64) ([]WalletHistory, error)
	LogTransaction(userID int64, amount int64, txType string) error
}

type WalletUsecase interface {
	GetBalance(userID int64) (*Wallet, error)
	Withdraw(userID int64, amount int64) error
	Deposit(userID int64, amount int64) error
	GetHistory(userID int64) ([]WalletHistory, error)
}
