package utils

import (
	"database/sql"

	"github.com/RipulHandoo/blogx/authentication/db/database"
	"github.com/google/uuid"
)

type DbUserFullSchema struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func MapRegisteredUser(DbUser database.User) DbUserFullSchema{
	return DbUserFullSchema{
		ID:        DbUser.ID,
		FirstName: DbUser.FirstName,
		LastName:  DbUser.LastName,
		Email:     DbUser.Email,
		Bio:       DbUser.Bio.String,
		CreatedAt: DbUser.CreatedAt.String(),
		UpdatedAt: DbUser.UpdatedAt.String(),
	}
}

type DBUserResponse struct {
	ID        uuid.UUID      `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Bio       sql.NullString `json:"bio"`
}

func MapLoginUser(dbUser database.User) DBUserResponse {
	return DBUserResponse{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Bio:       dbUser.Bio,
	}
}