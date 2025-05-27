// filename: webhook_dump_server.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s  %s  %s", r.RemoteAddr, r.Method, r.URL.RequestURI())

		// print all headers
		for k, v := range r.Header {
			log.Printf("Header[%q] = %q", k, v)
		}

		// print body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("read body error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		log.Printf("---- BODY BEGIN ----\n%s\n---- BODY  END  ----", body)

		// send response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	log.Println("http dump server listening on :8080 …")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
