package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", GetParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", Bad).Methods(http.MethodGet)
	router.HandleFunc("/data", SetParam).Methods(http.MethodPost)
	router.HandleFunc("/headers", ReadHeader).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

// main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func GetParam(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// param := params["PARAM"]
	param := strings.TrimPrefix(r.URL.Path, "/name/") // alternative
	msg := fmt.Sprintf("Hello, %s!", param)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func Bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func SetParam(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("body read error %s", err), http.StatusBadRequest)
		return
	}
	msg := fmt.Sprintf("I got message:\n%s", reqBody)
	w.Write([]byte(msg))
}

func ReadHeader(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	numA, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid input %s", err), http.StatusBadRequest)
		return
	}
	numB, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid input %s", err), http.StatusBadRequest)
		return
	}
	sum := numA + numB
	result := strconv.Itoa(sum)
	w.Header().Set("a+b", result)
	w.WriteHeader(http.StatusOK)
}
