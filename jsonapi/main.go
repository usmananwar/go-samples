package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	var router = mux.NewRouter()
	router.HandleFunc("/hello", greetingsHandler).Methods("GET")
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	router.HandleFunc("/queryMsg", handleQueryMsgs).Methods("GET")
	router.HandleFunc("/urlMsg/{msg}", handleURLMsgs).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
func greetingsHandler(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Println("Hello world")
}

func healthCheckHandler(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	json.NewEncoder(responseWriter).Encode("Still Alive!")
}

func handleQueryMsgs(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	vars := httpRequest.URL.Query()
	message := vars.Get("msg")
	json.NewEncoder(responseWriter).Encode(map[string]string{"Received Message": message})

}

func handleURLMsgs(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	message := vars["msg"]
	json.NewEncoder(responseWriter).Encode(map[string]string{"Received Message": message})

}
