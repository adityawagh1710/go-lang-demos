package repository

import (
	"context"
	"project/internal/db"
	"project/internal/model"
)

func GetUserByID(ctx context.Context, id int64) (model.User, error) {
	var user model.User

	query := `SELECT id, name, email FROM users WHERE id=$1`

	err := db.DB.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Name, &user.Email)

	return user, err
}
