package repository

import (
	"database/sql"
	"fmt"

	"github.com/MoodyShoo/GinAPI/internal/models"
)

type UserRepository interface {
	GetUser(id uint) (*models.User, error)
	InsertUser(u models.User) (int64, error)
	DeleteTestUsers() (int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUser(id uint) (*models.User, error) {
	query := `SELECT id, login, name, gender, age, phone, email, avatar, register_time, is_active 
		FROM users 
		WHERE id = $1`

	row := r.db.QueryRow(query, id)

	user := new(models.User)

	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Name,
		&user.Gender,
		&user.Age,
		&user.Contacts.Phone,
		&user.Contacts.Email,
		&user.Avatar,
		&user.RegisterTime,
		&user.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("пользователь не найден")
		}
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	return user, nil
}

func (r *userRepository) InsertUser(u models.User) (int64, error) {
	res, err := r.db.Exec(`
		INSERT INTO users (login, name, gender, age, phone, email)
		VALUES ($1,$2,$3,$4,$5,$6)`,
		u.Login, u.Name, u.Gender, u.Age, u.Contacts.Phone, u.Contacts.Email)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *userRepository) DeleteTestUsers() (int, error) {
	res, err := r.db.Exec(`DELETE FROM users WHERE login LIKE 'user%'`)
	if err != nil {
		return 0, err
	}
	n, _ := res.RowsAffected()
	return int(n), nil
}
