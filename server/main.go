package main

import (
	"fmt"
	"net/http"
)

func GetEntriesByKey(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Printf("GET coming\n")
}

func PutEntriesByKey(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Printf("PUT coming\n")
}

func HandleEntries(
	w http.ResponseWriter,
	r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetEntriesByKey(w, r)
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
