package database

import (
	"database/sql"
	"fmt"

	repository "github.com/MoodyShoo/GinAPI/internal/database/repository/user_repository"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
	us repository.UserRepository
}

func NewDatabase() (*Database, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к БД: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при проверке подключения: %w", err)
	}

	if err := createUsersTable(db); err != nil {
		return nil, fmt.Errorf("ошибка при создании таблицы: %w", err)
	}

	return &Database{
		DB: db,
		us: repository.NewUserRepository(db),
	}, nil

}

func createUsersTable(db *sql.DB) error {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        login VARCHAR(255) NOT NULL,
        name VARCHAR(255) NOT NULL,
        gender INTEGER NOT NULL,
        age INTEGER,
        phone VARCHAR(20),
        email VARCHAR(255),
        avatar BYTEA,
        register_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        is_active BOOLEAN DEFAULT TRUE
    );
    `

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы users: %w", err)
	}

	return nil
}

func (d *Database) Close() {
	d.DB.Close()
}

func (d *Database) US() repository.UserRepository {
	return d.us
}
