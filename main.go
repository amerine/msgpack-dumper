package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
		dec := decoder.NewDecoder(rdr)

		count := 0
		for {
			ret, _, record := decoder.GetRecord(dec)
			if ret != 0 {
				break
			}

			timestamp := time.Now()
			fmt.Printf("[%d] %s: [%s, {", count, "empty", timestamp.String())
			for k, v := range record {
				fmt.Printf("\"%s\": %v, ", k, v)
			}
			fmt.Printf("}]\n")
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
