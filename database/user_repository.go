package database

import (
	"database/sql"
	"sarkaz_api/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

func (repo *UserRepository) GetUsersById(id int) (models.User, error) {
	var user models.User
	error := repo.db.QueryRow("SELECT * FROM Users WHERE userId=?", id).Scan(
		&user.UserId, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.LastLogin,
	)
	if error != nil {
		return user, error
	}
	return user, nil
}

func (repo *UserRepository) GetUsersByName(name string) (models.User, error) {
	var user models.User
	error := repo.db.QueryRow("SELECT * FROM Users WHERE username=?", name).Scan(
		&user.UserId, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.LastLogin,
	)
	if error != nil {
		return user, error
	}
	return user, nil
}

func (repo *UserRepository) GetUserProfile(id int) (models.UserProfile, error) {
	var user models.UserProfile
	error := repo.db.QueryRow("SELECT * FROM Users WHERE userId=?", id).Scan(
		&user.UserId, &user.Username, &user.Bio, &user.Signature, &user.FavCharId,
	)
	if error != nil {
		return user, error
	}
	return user, nil
}