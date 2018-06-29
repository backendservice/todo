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
	instance := NewInMemoryService()
	http.HandleFunc("/service/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			result := instance.List()
			json.NewEncoder(w).Encode(result) // TODO: implement
		case "PUT":
			fallthrough
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
			if r.Method == "POST" {
				instance.Create(item)
				w.WriteHeader(http.StatusCreated)
			} else {
				if err := instance.Update(item); err != nil {
					if serviceErr, ok := err.(*ServiceError); ok {
						w.WriteHeader(serviceErr.Code)
					} else {
						w.WriteHeader(http.StatusInternalServerError)
					}
					return
				}
				w.WriteHeader(http.StatusOK)
			}
		case "DELETE":
			id := getID(r)
			instance.Delete(id)
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
