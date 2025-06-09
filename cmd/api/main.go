package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/xid"

	"github.com/clock-en/golang-grpc-tutorial/internal/greet"
)

func main() {
	xid := xid.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logText := fmt.Sprintf("XID: %s", xid)
		log.Println(logText)
		fmt.Fprintf(w, greet.Hello())
	})

	http.ListenAndServe(":8080", nil)
}
