package repositories

import (
	"database/sql"

	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type RefreshTokensRepository interface {
	FindOne(refreshToken string) (*models.RefreshToken, error)
	Create(refreshToken string, userID int) (*models.RefreshToken, error)
	Delete(refreshToken string) error
}

type refreshTokens struct {
	db *sql.DB
}

func NewRefreshTokens(db *sql.DB) RefreshTokensRepository {
	return &refreshTokens{db}
}

func (r *refreshTokens) FindOne(refreshToken string) (*models.RefreshToken, error) {
	row := r.db.QueryRow(`
		SELECT user_id
		FROM 
			refresh_tokens
		WHERE token = $1`,
		refreshToken,
	)
	return scanRefreshToken(row)
}

func (r *refreshTokens) Create(refreshToken string, userID int) (*models.RefreshToken, error) {
	row := r.db.QueryRow(`
		INSERT INTO refresh_tokens 
			(token, user_id) 
		VALUES 
			($1, $2)
		RETURNING id, token`,
		refreshToken,
		userID,
	)
	return scanRefreshToken(row)
}

func (r *refreshTokens) Delete(refreshToken string) error {
	_, err := r.db.Exec(`
		DELETE FROM
			refresh_tokens
		WHERE token = $1`,
		refreshToken,
	)
	return err
}

func scanRefreshToken(row *sql.Row) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := row.Scan(&token.ID, &token.Token, &token.UserID)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
