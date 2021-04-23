package repository

import "blogapi/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
}
