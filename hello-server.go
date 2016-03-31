package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var port = flag.String("p", "5000", "port to listen on (default 5000)")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})

	fmt.Println("Listening on port:", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		fmt.Println(err)
		return
	}
}
