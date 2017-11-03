package main

import (
	"encoding/json"
	"log"
	"meander"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	meander.APIKey = "AIzaSyA7qw8gDZEePvIUKZbPOgXH-Hekz0eNNvY"
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	log.Println("サーバを起動します。")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

// データをレスポンスに書き出す処理を抽象化
func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
