package DataBase

import (
	Usr "../Models"
	L "../utils/Logger"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"os"
	"time"
)

var ctx = context.Background()

var client *redis.Client

func CheckConnect() bool  {
	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("dbAdr")+":"+os.Getenv("dbPort"),
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping(client.Context()).Result()
	L.ErorrLog(err)
	if pong != "" {
		return true
	}else {
		return false
	}

}

func SetValues(K string,Value string,dur int)  {
	// we can call set with a `Key` and a `Value`.
	err := client.Set(ctx, K, Value,time.Duration(dur) * time.Minute).Err()
	// if there has been an error setting the value
	// handle the error
	L.ErorrLog(err)

}

func GetAuthStatus() *Usr.User {
	val, err := client.Get(ctx, "Auth").Result()
	L.ErorrLog(err)
	//L.ErorrDebugger( err)
	var User Usr.User
	json.Unmarshal([]byte(val), &User)
	return &User
}

func SetAuth(user *Usr.User)  {
	user.Status=true

	b, err := json.Marshal(user)


	err = client.Set(ctx, "Auth", b,time.Duration(1) * time.Hour).Err()

	L.ErorrLog(err)

}
