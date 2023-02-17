package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-example/models"

	"github.com/go-redis/redis/v8"
)

func CreateUser(name string, age uint, phone string, email string, address string) error {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	user := models.User{
		Name:    name,
		Age:     age,
		Phone:   phone,
		Email:   email,
		Address: address,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = redisClient.LPush(ctx, "create_user_list", userJson).Err()
	if err != nil {
		return err
	}

	fmt.Println("CreateNewUser task added to the Redis queue")
	return nil
}
