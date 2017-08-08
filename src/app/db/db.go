package db

import "app/models"

type DB interface {
	createUser(user *models.User)
	getUsers(users *[]models.User) *[]models.User
}
