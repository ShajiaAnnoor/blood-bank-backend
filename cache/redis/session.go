package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"gitlab.com/Aubichol/blood-bank-backend/cache"
	"gitlab.com/Aubichol/blood-bank-backend/model"
)

const (
	sessionPrefix string = "session"
)

type session struct {
	c *redis.Client
}

func (s *session) Create(data *model.Session) error {
	data.Key = fmt.Sprintf("%s_%s_%s", sessionPrefix, data.UserID, uuid.New().String())
	duration := data.ExpiredAt.Sub(data.CreatedAt)
	bytes, err := data.ToByte()

	if err != nil {
		return err
	}

	result := s.c.Set(data.Key, bytes, duration)
	return result.Err()
}

func (s *session) GetByKey(key string) (*model.Session, error) {
	result := s.c.Get(key)
	if result.Err() == redis.Nil {
		return nil, nil
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	data := model.Session{}
	bytes, err := result.Bytes()
	if err != nil {
		return nil, err
	}

	if err = data.FromBytes(bytes); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *session) GetByUserID(userID string) ([]*model.Session, error) {
	return nil, nil
}

func (s *session) RemoveByKey(id string) error {
	return nil
}

func NewSession(c *redis.Client) cache.Session {
	return &session{c}
}
