package handler

import (
	"backend/api/storer"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Service interface {
	CreateSeed() map[string]string
	DecodeSeed(seed string) map[string]string
	EncodeSeed(m map[string]string) string
}

// ToDo: доделать все методы
type Storer interface {
	CreateSession(ctx context.Context, s *storer.Session, adminID string) (*storer.Session, error)
	CreateAccount(ctx context.Context, a *storer.Account) (*storer.Account, error)
	CreatePlayerCard(ctx context.Context, pc *storer.PlayerCard, sessionID string, accountID string) (*storer.PlayerCard, error)
	ListSessionByID(ctx context.Context, account_id string) ([]storer.Session, error)
}

type handler struct {
	storer  Storer
	service Service
	context context.Context
}

func New(str Storer, src Service) *handler {
	return &handler{
		storer:  str,
		service: src,
		context: context.Background(),
	}
}

func (h *handler) generateSeed(w http.ResponseWriter, r *http.Request) error {
	seed := h.service.CreateSeed()
	// encodedSeed := h.service.EncodeSeed(seed)

	return writeResponse(w, seed, http.StatusOK)
}

func (h *handler) handleListSessionByID(w http.ResponseWriter, r *http.Request) error {
	accountID := r.Header.Get("id")

	sessions, err := h.storer.ListSessionByID(h.context, accountID)
	if err != nil {
		return err
	}

	return writeResponse(w, sessions, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, res any, status int) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(res)
}

type APIfunc func(w http.ResponseWriter, r *http.Request) error

func handleFuncWrapper(f APIfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		}
	}
}
