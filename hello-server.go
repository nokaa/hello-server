package main

import (
	"fmt"
	"net/http"
)

const PORT = ":5000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})

	fmt.Println("Listening on port:", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Println(err)
		return
	}
}
