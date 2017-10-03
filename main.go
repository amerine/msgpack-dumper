package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vmihailenco/msgpack"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		if r.Header.Get("Content-Type") != "msgpack" {
			http.Error(w, "Invalid Content-Type: "+r.Header.Get("Content-Type"), http.StatusBadRequest)
		}

		defer r.Body.Close()
		dec := msgpack.NewDecoder(r.Body)

		out, err := dec.DecodeMap()
		if err != nil {
			http.Error(w, "Invalid Payload", http.StatusBadRequest)
		}

		fmt.Printf("msg: [%+v]", out)
		w.WriteHeader(http.StatusNoContent)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
