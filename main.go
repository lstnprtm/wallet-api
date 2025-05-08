package main

import (
	"fmt"
	"os"

	"github.com/lstnprtm/wallet-api/config"
	"github.com/lstnprtm/wallet-api/internal/handler"
	"github.com/lstnprtm/wallet-api/internal/repository"
	"github.com/lstnprtm/wallet-api/internal/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_ = godotenv.Load()
	conf := config.LoadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Name)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Auth handler
	authRepo := repository.NewAuthRepo(db)
	authUC := usecase.NewAuthUsecase(authRepo)
	handler.NewAuthHandler(e, authUC)

	// Wallet handler
	walletRepo := repository.NewWalletRepo(db)
	walletUC := usecase.NewWalletUsecase(walletRepo)

	api := e.Group("/api")

	api.Use(JWTMiddleware())
	handler.NewWalletHandler(api, walletUC)

	e.Logger.Fatal(e.Start(":8080"))
}

func JWTMiddleware() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET")
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		ContextKey: "user_id",
	})
}
