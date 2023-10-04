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
	query := "SELECT * FROM Users WHERE userId=?"

	stmt, error := repo.db.Prepare(query)
	if error != nil {
		return user, error
	}
	defer stmt.Close()

	err := stmt.QueryRow(id).Scan(
		&user.UserId, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.LastLogin,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) GetUsersByName(name string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM Users WHERE username=?"

	stmt, error := repo.db.Prepare(query)
	if error != nil {
		return user,error
	}
	defer stmt.Close()

	err := stmt.QueryRow(name).Scan(
		&user.UserId, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.LastLogin,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserProfile(id int) (models.UserProfile, error) {
	var user models.UserProfile
	query := "SELECT * FROM Users WHERE userId=?"

	stmt, error := repo.db.Prepare(query)
	if error != nil {
		return user, error
	}
	defer stmt.Close()

	err := repo.db.QueryRow("SELECT * FROM Users WHERE userId=?", id).Scan(
		&user.UserId, &user.Username, &user.Bio, &user.Signature, &user.FavCharId,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}