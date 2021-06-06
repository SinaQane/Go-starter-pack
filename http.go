package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
Number int       `json:"number"`
Data   string    `json:"data"`
Time   time.Time `json:"time"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("hello, world!"))
}

func echo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(body)
}

func jsonEncoder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // value is the Content-Type with is json here
	encoder := json.NewEncoder(w)
	_ := encoder.Encode(Response{
		Number: 1000,
		Data:   "data",
		Time:   time.Now(),
	})
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/json", jsonEncoder)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}