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
	instance := NewSimpleDBService("usol")
	http.HandleFunc("/service/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			items, err := instance.List();
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(items) // TODO: implement
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
				if err := instance.Upsert(item, false); err == nil{
					w.WriteHeader(http.StatusCreated)
				}else{
					http.Error(w, err.Error(), 500)
					return
				}
			}else{
				if err := instance.Upsert(item, true); err != nil{
					if svcErr, ok := err.(*ServiceError); ok {
						w.WriteHeader(svcErr.Code)
					}else{
						w.WriteHeader(http.StatusOK)
					}
				return
				}
			}
		case "DELETE":
			if err := instance.Delete(getID(r)); err != nil{
				http.Error(w, err.Error(), 500)
			} else {
				w.WriteHeader(http.StatusOK)
			}
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
