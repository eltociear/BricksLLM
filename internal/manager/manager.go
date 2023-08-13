package manager

import (
	"time"

	"github.com/bricks-cloud/bricksllm/internal/key"
	"github.com/bricks-cloud/bricksllm/internal/util"
)

type Storage interface {
	GetKeysByTag(tag string) ([]*key.ResponseKey, error)
	UpdateKey(id string, key *key.UpdateKey) (*key.ResponseKey, error)
	CreateKey(key *key.RequestKey) (*key.ResponseKey, error)
	DeleteKey(id string) error
}

type Manager struct {
	s Storage
}

func NewManager(s Storage) *Manager {
	return &Manager{
		s: s,
	}
}

func (m *Manager) GetKeysByTag(tag string) ([]*key.ResponseKey, error) {
	return m.s.GetKeysByTag(tag)
}

func (m *Manager) CreateKey(rk *key.RequestKey) (*key.ResponseKey, error) {
	rk.CreatedAt = time.Now().Unix()
	rk.UpdatedAt = time.Now().Unix()
	rk.KeyId = util.NewUuid()
	f := false
	rk.Revoked = &f

	if err := rk.Validate(); err != nil {
		return nil, err
	}

	return m.s.CreateKey(rk)
}

func (m *Manager) UpdateKey(id string, uk *key.UpdateKey) (*key.ResponseKey, error) {
	uk.UpdatedAt = time.Now().Unix()

	if err := uk.Validate(); err != nil {
		return nil, err
	}

	return m.s.UpdateKey(id, uk)
}

func (m *Manager) DeleteKey(id string) error {
	return m.s.DeleteKey(id)
}
