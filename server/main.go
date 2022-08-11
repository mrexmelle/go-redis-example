package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
)

type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Value struct {
	Value string `json:"value"`
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
	ctx context.Context,
	rdb *redis.Client,
	key string,
	value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

func HandleEntries(
	w http.ResponseWriter,
	r *http.Request) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})

	vars := mux.Vars(r)
	key, _ := vars["key"]

	switch r.Method {
	case http.MethodGet:
		value := GetEntriesByKey(
			r.Context(),
			rdb,
			key,
		)

		response, _ := json.Marshal(value)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")
		w.Write(response)

		break
	case http.MethodPut:
		err := PutEntriesByKey(
			r.Context(),
			rdb,
			key,
			"matt and jess",
		)
		if err == nil {
			response, _ := json.Marshal(Entry{key, "matt and jess"})
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-type", "application/json")
			w.Write(response)
		} else {
			http.Error(w, "Redis not accessible", http.StatusServiceUnavailable)
		}
		break
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		break
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/entries/{key}", HandleEntries)
	http.ListenAndServe(":8080", router)
}
