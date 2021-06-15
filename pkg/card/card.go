package card

import (
	"context"
	"sync"
)

type Card struct {
	Id       int64
	Number   string
	Issuer   string
	Owner    Owner
	NameCard string
	Type     string
}

type Owner struct {
	Name     string
	Lastname string
}

type Service struct {
	mu    sync.RWMutex
	Cards []*Card
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) All(context.Context) []*Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Cards
}
