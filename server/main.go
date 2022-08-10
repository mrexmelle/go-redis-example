package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v9"
)

func GetEntriesByKey(
	ctx context.Context,
	rdb *redis.Client,
	key string) {

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("GET err\n")
	} else {
		fmt.Printf("GET: %s\n", val)
	}
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
		GetEntriesByKey(
			r.Context(),
			rdb,
			"abc",
		)
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
