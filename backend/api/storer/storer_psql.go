package storer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgreSQLStorer struct{ db *sqlx.DB }

func NewPostgreSQLStorer(db *sqlx.DB) *PostgreSQLStorer {
	return &PostgreSQLStorer{
		db: db,
	}
}

func (ps *PostgreSQLStorer) CreateAccount(ctx context.Context, a *Account) (*Account, error) {
	var accountID string
	q := "insert into account (email, password, created_at) values (:id, :email, :password) returning id"

	err := ps.db.GetContext(ctx, &accountID, q, a)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}

	a.ID = accountID

	return a, nil
}

func (ps *PostgreSQLStorer) ListAccount(ctx context.Context) ([]Account, error) {
	var accounts []Account
	q := "select * from account"

	if err := ps.db.SelectContext(ctx, accounts, q); err != nil {
		return nil, fmt.Errorf("error listing accounts: %w", err)
	}

	return accounts, nil
}

func (ps *PostgreSQLStorer) UpdateAccount(ctx context.Context, a *Account) (*Account, error) {
	q := "update account set email=:email, password=:password "

	_, err := ps.db.NamedExecContext(ctx, q, a)
	if err != nil {
		return nil, fmt.Errorf("error updating account: %w", err)
	}

	return a, nil
}

func (ps *PostgreSQLStorer) DeleteAccount(ctx context.Context, id string) error {
	q := "delete from account where id=?"

	_, err := ps.db.NamedExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}

	return nil
}

func (ps *PostgreSQLStorer) CreateSession(ctx context.Context, s *Session, adminID string) (*Session, error) {
	var sessionID string
	s.AdminID = adminID
	q := "insert into session (name, admin_id) values (:name, :admin_id) returning id"

	err := ps.db.GetContext(ctx, &sessionID, q, s)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}

	s.ID = sessionID

	return s, nil
}

func (ps *PostgreSQLStorer) UpdateSession(ctx context.Context, s *Session) (*Session, error) {
	q := "update session set name=:name, seed=:seed, picture=:picture"

	_, err := ps.db.NamedExecContext(ctx, q, s)
	if err != nil {
		return nil, fmt.Errorf("error updating seesion: %w", err)
	}

	return s, nil
}

func (ps *PostgreSQLStorer) UpdateSessionSeed(ctx context.Context, seed string) error {
	q := "update session set seed=?"

	_, err := ps.db.NamedExecContext(ctx, q, seed)
	if err != nil {
		return fmt.Errorf("error adding seed to session: %w", err)
	}

	return nil
}

func (ps *PostgreSQLStorer) DeleteSession(ctx context.Context, id string) error {
	q := "delete from session where id=?"

	_, err := ps.db.NamedExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}

	return nil
}

func (ps *PostgreSQLStorer) ListSession(ctx context.Context) ([]Session, error) {
	var sessions []Session
	q := "select * from session"

	err := ps.db.SelectContext(ctx, sessions, q)
	if err != nil {
		return nil, fmt.Errorf("error listing sessions: %w", err)
	}

	return sessions, nil
}

func (ps *PostgreSQLStorer) ListSessionByID(ctx context.Context, account_id string) ([]Session, error) {
	var sessions []Session
	q := `SELECT s.*, (SELECT COUNT(*) FROM "player_card" pc WHERE s.id = pc.session_id) AS player_count FROM "session" s WHERE EXISTS (SELECT 1 FROM "player_card" pc WHERE s.id = pc.session_id AND pc.account_id=$1);`

	err := ps.db.SelectContext(ctx, &sessions, q, account_id)
	if err != nil {
		return nil, fmt.Errorf("error listing sessions by id: %w", err)
	}

	return sessions, nil
}

func (ps *PostgreSQLStorer) CreatePlayerCard(ctx context.Context, pc *PlayerCard, sessionID string, accountID string) (*PlayerCard, error) {
	var playerCardID int64

	pc.SessionID = sessionID
	pc.AccountID = accountID
	q := `insert into session (account_id, session_id, nickname, preferences) values (:account_id, :session_id, :nickname, :preferences) returning id`

	err := ps.db.GetContext(ctx, &playerCardID, q, pc)
	if err != nil {
		return nil, fmt.Errorf("error creating account session: %w", err)
	}

	pc.ID = playerCardID

	return pc, nil
}

func (ps *PostgreSQLStorer) ListPlayerCardByID(ctx context.Context, pc *PlayerCard, accountID string) ([]PlayerCard, error) {
	var playerCards []PlayerCard
	q := "select * from player_card where id=?"

	err := ps.db.SelectContext(ctx, playerCards, q, accountID)
	if err != nil {
		return nil, fmt.Errorf("error listing player cards by id: %w", err)
	}

	return playerCards, nil
}
