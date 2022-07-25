package repository

import (
	"log"

	models "github.com/miguelanselmo/my-web-app/models"
	utils "github.com/miguelanselmo/my-web-app/utils"
)

var userRepo = map[string]models.UserModel{}

type RepositorySrvc struct {
	log *log.Logger
}

func New(log *log.Logger) *RepositorySrvc {
	userRepo = map[string]models.UserModel{}
	return &RepositorySrvc{log: log}
}

type UserRepositoryImpl interface {
	GetUser(email string) (models.UserModel, error)
	CreateUser(user models.UserModel) error
}

func (*RepositorySrvc) CreateUser(user models.UserModel) {
	passwordHash, err := utils.GetPasswordHash(user.Password)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	userRepo[user.Email] = models.UserModel{
		Email:    user.Email,
		Password: passwordHash,
	}
	log.Println("User created")
	log.Println("DB: ", userRepo)
}

func (*RepositorySrvc) GetUser(user models.UserModel) (models.UserModel, bool) {
	user, ok := userRepo[user.Email]
	return user, ok
}
