package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	r "gopkg.in/gorethink/gorethink.v3"
)

var (
	router  *mux.Router
	session *r.Session
)

const (
	entry  = "./client/dist/index.html"
	static = "./client/dist"
)

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "wespk",
		MaxOpen:  40,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func NewServer(port string) *http.Server {
	router = initRouting()

	return &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

}

func StartServer(server *http.Server) {
	log.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error: %v", err)
	}
}

func initRouting() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/todos", GetTodoHandler).Methods("GET")
	api.HandleFunc("/todos/new", NewTodoHandler)

	r.PathPrefix("/static").Handler(http.FileServer(http.Dir(static)))
	r.PathPrefix("/").HandlerFunc(indexHandler(entry))

	return r
}

func indexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

func GetTodoHandler(w http.ResponseWriter, req *http.Request) {
	items := []TodoItem{}

	res, err := r.Table("todos").OrderBy(r.Asc("Created")).Run(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = res.All(&items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&items)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write(b)
}

func NewTodoHandler(w http.ResponseWriter, req *http.Request) {
	item := NewTodoItem("New Todo", "New Project")
	item.Created = time.Now()

	_, err := r.Table("todos").Insert(item).RunWrite(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
