package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	addr = ":50080"
)

func main() {
	instance := NewSimpleDBService("chloe")
	http.HandleFunc("/service/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if result, err := instance.List(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(result)
			}
		case "POST":
			var item Item
			if r.Body == nil {
				http.Error(w, "Please send a request body", 400)
				return
			}
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			if err := instance.Create(item); err != nil {
				http.Error(w, "InternalServerError", 500)
			}
			w.WriteHeader(http.StatusCreated)

		case "PUT":
			var item Item
			if r.Body == nil {
				http.Error(w, "Please send a request body", 400)
				return
			}
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			instance.Update(item)
			w.WriteHeader(http.StatusOK)

		case "DELETE":
			id := getID(r)
			instance.Delete(id)
			w.WriteHeader(http.StatusOK)
		}
	})

	handleStatic()

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func handleStatic() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/todo/", http.StripPrefix("/todo/", fs))
}

func getID(r *http.Request) string {
	pathElems := strings.Split(r.RequestURI, "/")
	return pathElems[len(pathElems)-1]
}
