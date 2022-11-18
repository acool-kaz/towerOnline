package service

import (
	"errors"

	"github.com/acool-kaz/towerOnline/internal/models"
	"github.com/acool-kaz/towerOnline/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(user models.User) error
	GetOne(telegramId int64) (models.User, error)
}

type UserService struct {
	stor storage.User
}

func newUserService(stor storage.User) *UserService {
	return &UserService{
		stor: stor,
	}
}

func (s *UserService) Create(user models.User) error {
	if _, err := s.stor.GetOne(user.TelegramId); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}
	}
	return s.stor.Create(user)
}

func (s *UserService) GetOne(telegramId int64) (models.User, error) {
	return s.stor.GetOne(telegramId)
}
