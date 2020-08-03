package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	fmt.Printf("Listening on : %+v", port)
	http.ListenAndServe(":"+port, r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Docker test server running!")

}
