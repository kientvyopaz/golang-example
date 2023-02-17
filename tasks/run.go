package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func run_queue() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for {
		result, err := rdb.BRPop(ctx, 0, "create_user_list").Result()
		if err != nil {
			fmt.Println(err)
			return
		}

		var user User
		err = json.Unmarshal([]byte(result[1]), &user)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Processing user %d: %s (%s)\n", user.ID, user.Name, user.Email)
	}
}
