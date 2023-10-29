package repositories

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type RefreshTokensRepository interface {
	FindOne(refreshToken string) (*models.RefreshToken, error)
	Upsert(refreshToken string, userID int) (*models.RefreshToken, error)
	Delete(refreshToken string) error
	DeleteByUserId(userID int) error
}

type refreshTokens struct {
	db *sql.DB
}

func NewRefreshTokens(db *sql.DB) RefreshTokensRepository {
	return &refreshTokens{db}
}

func (r *refreshTokens) FindOne(refreshToken string) (*models.RefreshToken, error) {
	row := r.db.QueryRow(`
		SELECT id, token, user_id
		FROM 
			refreshtokens
		WHERE token = $1`,
		refreshToken,
	)
	return scanRefreshToken(row)
}

func (r *refreshTokens) Upsert(refreshToken string, userID int) (*models.RefreshToken, error) {
	err := r.DeleteByUserId(userID)
	if err != nil {
		log.Warn(err)
	}
	row := r.db.QueryRow(`
		INSERT INTO refreshtokens 
			(token, user_id) 
		VALUES 
			($1, $2)
		RETURNING id, token, user_id`,
		refreshToken,
		userID,
	)
	return scanRefreshToken(row)
}

func (r *refreshTokens) Delete(refreshToken string) error {
	_, err := r.db.Exec(`
		DELETE FROM
			refreshtokens
		WHERE token = $1`,
		refreshToken,
	)
	return err
}

func (r *refreshTokens) DeleteByUserId(userID int) error {
	_, err := r.db.Exec(`
		DELETE FROM
			refreshtokens
		WHERE user_id = $1`,
		userID,
	)
	return err
}

func scanRefreshToken(row *sql.Row) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := row.Scan(&token.ID, &token.Token, &token.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFoundError{Message: "token not found"}
		}
		return nil, err
	}
	return &token, nil
}
