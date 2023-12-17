package main

import (
	"net/http"

	"github.com/dickymuliafiqri/BenchBox/server/api/bench"
)

func main() {
	http.HandleFunc("/bench", bench.PostBench)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
