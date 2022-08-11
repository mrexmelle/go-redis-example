package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v9"
)

type Value struct {
	Value string
}

func GetEntriesByKey(
	ctx context.Context,
	rdb *redis.Client,
	key string) Value {

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		val = ""
	}

	return Value{val}
}

func PutEntriesByKey(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Printf("PUT coming\n")
}

func HandleEntries(
	w http.ResponseWriter,
	r *http.Request) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})

	switch r.Method {
	case http.MethodGet:
		value := GetEntriesByKey(
			r.Context(),
			rdb,
			"abc",
		)

		response, _ := json.Marshal(value)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")
		w.Write(response)

		break
	case http.MethodPut:
		PutEntriesByKey(w, r)
		break
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		break
	}
}

func main() {
	http.HandleFunc("/entries", HandleEntries)
	http.ListenAndServe(":8080", nil)
}
