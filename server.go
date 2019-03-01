package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8001", "")
	flag.Parse()

	fs := http.FileServer(http.Dir("lighthouse"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
