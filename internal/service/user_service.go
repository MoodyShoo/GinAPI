package service

import (
	"fmt"

	repository "github.com/MoodyShoo/GinAPI/internal/database/repository/user_repository"
	"github.com/MoodyShoo/GinAPI/internal/models"
)

type UserService interface {
	GetUser(id uint) (*models.User, error)
	GetUserField(id uint, field string) (interface{}, error)
	SeedTestUsers() (int, error)
	EraseTestUsers() (int, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *userService) GetUserField(id uint, field string) (interface{}, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	switch field {
	case "login":
		return user.Login, nil
	case "name":
		return user.Name, nil
	case "age":
		return user.Age, nil
	case "email":
		return user.Contacts.Email, nil
	case "phone":
		return user.Contacts.Phone, nil
	case "gender":
		return user.Gender, nil
	case "is_active":
		return user.IsActive, nil
	default:
		return nil, fmt.Errorf("неизвестное поле: %s", field)
	}
}

func (s *userService) SeedTestUsers() (int, error) {
	users := []models.User{
		{Login: "user1", Name: "Alice", Gender: models.Female, Age: 25, Contacts: models.Contacts{Phone: "123", Email: "alice@example.com"}},
		{Login: "user2", Name: "Bob", Gender: models.Male, Age: 30, Contacts: models.Contacts{Phone: "456", Email: "bob@example.com"}},
	}

	count := 0
	for _, u := range users {
		_, err := s.repo.InsertUser(u)
		if err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (s *userService) EraseTestUsers() (int, error) {
	return s.repo.DeleteTestUsers()
}
