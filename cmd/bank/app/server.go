package app

import (
	"encoding/json"
	"github.com/ArtDark/bgo_server/cmd/bank/app/dto"
	"github.com/ArtDark/bgo_server/pkg/card"
	"log"
	"net/http"
)

// Server для обработки сервиса картgo
type Server struct {
	cardSvc *card.Service
	mux     *http.ServeMux
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux) *Server {
	return &Server{cardSvc: cardSvc, mux: mux}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/addCard", s.addCard)
	s.mux.HandleFunc("/editCard", s.editCard)
	s.mux.HandleFunc("/removeCard", s.removeCard)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {
	userCards := []*card.Card{}
	card := &card.Card{}

	err := json.NewDecoder(r.Body).Decode(card)

	for _, c := range s.cardSvc.All(r.Context()) {
		if c.Id == card.Id {
			userCards = append(userCards, c)
		}

	}

	dtos := make([]*dto.CardDTO, len(userCards))
	for i, c := range userCards {
		dtos[i] = &dto.CardDTO{
			Id:       c.Id,
			Number:   c.Number,
			Issuer:   c.Issuer,
			Owner:    c.Owner,
			NameCard: c.NameCard,
			Type:     c.Type,
		}
	}

	// TODO: вынести в отдельную функцию
	respBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	// по умолчанию статус 200 Ok
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) addCard(w http.ResponseWriter, r *http.Request) {
	cards := s.cardSvc.All(r.Context())
	card := &card.Card{
		Id:     s.cardSvc.CreateIdCard(),
		Number: "XXXX-XXXX-XXXX-XXXX",
		//TODO: можно реализовать через https://www.bincodes.com/api-creditcard-generator/
		Owner: s.cardSvc.GetOwner(),
	}

	err := json.NewDecoder(r.Body).Decode(card)
	if err != nil {
		log.Println(err)
		return
	}

	cards = append(cards, card)

	s.cardSvc.Cards = cards

}

func (s *Server) editCard(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}

func (s *Server) removeCard(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}
