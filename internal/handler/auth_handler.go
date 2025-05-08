package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lstnprtm/wallet-api/internal/domain"
	"net/http"
)

type AuthHandler struct {
	uc domain.AuthUsecase
}

func NewAuthHandler(e *echo.Echo, uc domain.AuthUsecase) {
	h := &AuthHandler{uc}
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	token, err := h.uc.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username and password required"})
	}
	if err := h.uc.Register(req.Username, req.Password); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "user registered"})
}
