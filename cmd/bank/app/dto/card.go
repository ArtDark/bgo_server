package dto

import "github.com/ArtDark/bgo_server/pkg/card"

type CardDTO struct {
	Id       int64      `json:"id"`
	Number   string     `json:"number"`
	Issuer   string     `json:"issuer"`
	Owner    card.Owner `json:"owner"`
	NameCard string     `json:"name_card"`
	Type     string     `json:"type"`
}
