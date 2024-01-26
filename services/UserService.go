package services

import (
	"backend_01/config"
	"backend_01/ent/user"
	"backend_01/middleware"
	"backend_01/models"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	client := config.SingletonClient()
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hashedPassword)
	createdUser, err := client.User.Create().SetUsername(user.Username).SetPassword(user.Password).SetRole(user.Role).Save(context.Background())
	if err != nil {
		panic(err)
	}
	log.Default().Println(createdUser)
}

func Login(a *models.User) (string, error) {
	// match hashed password with plain password
	GetFoods()

	client := config.UserClient()
	// find user by username
	u, err := client.Query().Where(user.Username(a.Username)).Only(context.Background())
	if err != nil {
		return "", err
	}
	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(a.Password))
	if err != nil {
		return "", err
	}
	// generate token
	token, err := middleware.GenerateToken(u)

	if err != nil {
		return "", err
	}
	return token, nil
}
