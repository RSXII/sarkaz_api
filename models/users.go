package models

type User struct {
	UserId 		int
	Username 	string
	FirstName	string
	LastName	string
	Email		string
	Password	string
	LastLogin	string
}

type UserProfile struct {
	UserId		int
	Username	string
	Bio 		string
	Signature	string
	FavCharId	int
}

type UserCollection struct {
	UserId			int
	CollectionId	int
}