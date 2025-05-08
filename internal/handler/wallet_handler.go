package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lstnprtm/wallet-api/internal/domain"
	"net/http"
)

type WalletHandler struct {
	uc domain.WalletUsecase
}

func NewWalletHandler(e *echo.Group, uc domain.WalletUsecase) {
	h := &WalletHandler{uc}
	e.GET("/balance", h.GetBalance)
	e.POST("/withdraw", h.Withdraw)
	e.POST("/deposit", h.Deposit)
	e.GET("/history", h.GetHistory)
}

func (h *WalletHandler) GetBalance(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	wallet, err := h.uc.GetBalance(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "not found"})
	}
	return c.JSON(http.StatusOK, wallet)
}

type amountReq struct {
	Amount int64 `json:"amount"`
}

func (h *WalletHandler) Withdraw(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	var req amountReq
	if err := c.Bind(&req); err != nil || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid amount"})
	}
	if err := h.uc.Withdraw(userID, req.Amount); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "withdraw successful"})
}

func (h *WalletHandler) Deposit(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	var req amountReq
	if err := c.Bind(&req); err != nil || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid amount"})
	}
	if err := h.uc.Deposit(userID, req.Amount); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "deposit successful"})
}

func (h *WalletHandler) GetHistory(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	history, err := h.uc.GetHistory(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, history)
}

func getUserIDFromToken(c echo.Context) (int64, error) {
	user := c.Get("user_id")
	token, ok := user.(*jwt.Token)
	if !ok {
		return 0, echo.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, echo.ErrUnauthorized
	}

	if uid, ok := claims["user_id"].(float64); ok {
		return int64(uid), nil
	}

	return 0, echo.ErrUnauthorized
}
