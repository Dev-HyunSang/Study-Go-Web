package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type User struct {
	FirstName string    `json: "FirstName"` // 강의에서는 First_Name으로 되었지만 오류 있어서 변경 First_Name => FirstName
	LastName  string    `json: "LastName"`
	Email     string    `json: "Email"`
	CreatedAt time.Time `json: "created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})
	return mux
}
