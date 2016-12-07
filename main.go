package main

import (
    "encoding/json"
    "log"
    "net/http"
    "math/rand"
    "github.com/gorilla/mux"
)

type Play struct {
    Play string `json:"play"`
}

type TableInfo struct {
    ID string `json:"id,omitempty"`
}

func TableUpdateEndpoint(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    plays := [4]string{"check", "fold", "call", "raise"}
    play := Play{Play: plays[rand.Intn(4)]}

    var tableInfo TableInfo
    _ = json.NewDecoder(req.Body).Decode(&tableInfo)

    json.NewEncoder(w).Encode(play)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/table-update", TableUpdateEndpoint).Methods("POST")

    log.Fatal(http.ListenAndServe(":4000", router))
}
