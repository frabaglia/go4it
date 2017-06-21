package main

import (
	"fmt"
	"net/http"
  "time"
  "log"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  // home := View{
  //   title: "Title of Go website /",
  //   body: "Body of this awsome Go site",
  // }

	fmt.Fprintf(w, "<h1>Home Go Lang! by GENOSHA</h1>")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
  // test := View{
  //   title: "Title of Go website /test",
  //   body: "Body of this awsome Go site",
  // }

	fmt.Fprintf(w, "<h1>Test Go Lang! by GENOSHA</h1>")
}

type View struct {
	title string
	body  string
}

func main() {

	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/test", testHandler)

  server := &http.Server{
		Addr:           port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Running GO server at " + port + "\n")
	log.Fatal(server.ListenAndServe())
}
