package main

import (
	"database/sql"
	"fmt"
	"github.com/lstnprtm/wallet-api/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dbConf := config.LoadEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB Connection failed: %v", err)
	}
	defer db.Close()

	runMigration(db)
	runSeeder(db)
}

func runMigration(db *sql.DB) {
	statements := []string{
		"DROP TABLE IF EXISTS wallet_histories;",
		"DROP TABLE IF EXISTS wallets;",
		"DROP TABLE IF EXISTS users;",

		`CREATE TABLE users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(50) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL
        );`,

		`CREATE TABLE wallets (
            id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT NOT NULL,
            balance BIGINT NOT NULL DEFAULT 0,
            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
        );`,

		`CREATE TABLE wallet_histories (
            id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT NOT NULL,
            amount BIGINT NOT NULL,
            type ENUM('withdraw', 'deposit') NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
        );`,
	}

	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			log.Fatalf("Migration error: %v", err)
		}
	}

	log.Println("✅ Migration completed.")
}

func runSeeder(db *sql.DB) {
	password := "secret123"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", "alice", string(hashed))
	if err != nil {
		log.Fatalf("Seeder user error: %v", err)
	}

	_, err = db.Exec("INSERT INTO wallets (user_id, balance) VALUES (?, ?)", 1, 150000)
	if err != nil {
		log.Fatalf("Seeder wallet error: %v", err)
	}

	log.Println("✅ Seeder completed. User: alice / Password:", password)
}
