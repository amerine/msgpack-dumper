package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/amerine/msgpack-dumper/decoder"
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

		if r.Header.Get("Content-Type") != "application/msgpack" {
			fmt.Println("Unsupported Content-Type Header: " + r.Header.Get("Content-Type"))
			http.Error(w, "Invalid Content-Type: "+r.Header.Get("Content-Type"), http.StatusBadRequest)
			return
		}

		dec := decoder.NewDecoder(r.Body)
		defer r.Body.Close()

		for {
			record, err := dec.GetRecord()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error Processing Record: " + err.Error())
				continue
			}

			fmt.Printf("\"record\": %q, ", record)
		}

		fmt.Println("Processed line")

	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
