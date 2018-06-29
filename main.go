package main

<<<<<<< HEAD
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
			json.NewEncoder(w).Encode([]Item{Item1()}) // TODO: implement
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
			instance.Create(item)
			w.WriteHeader(http.StatusCreated)
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
=======
import "fmt"

func main() {
	fmt.Println("Hello World")
>>>>>>> 1840a5d23656e7d9a6d0299cb7be6bef8da3ac0a
}
