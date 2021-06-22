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

func (s *Service) CreateIdCard() int64 {
	cards := s.Cards

	if len(cards) == 0 {
		return 0
	}
	lastNum := cards[len(cards)-1].Id

	return lastNum + 1

}
