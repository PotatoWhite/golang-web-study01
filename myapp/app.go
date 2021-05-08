package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" `
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	user := new(User)
	err := json.NewDecoder(request.Body).Decode(user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, "Bad Request : ", err)
		return
	}
	user.CreatedAt = time.Now().UTC()
	data, _ := json.Marshal(user)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	fmt.Fprint(writer, string(data))
}

func barHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	name := request.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	_, err := fmt.Fprintf(writer, "Hello %s", name)
	if err != nil {
		return
	}
}

func NewHttpHandler() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})

	return mux
}
