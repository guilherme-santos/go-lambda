package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

var supportedMethods = map[string]struct{}{
	"GET": struct{}{},
	// "POST":   struct{}{},
	// "PUT":    struct{}{},
	// "DELETE": struct{}{},
}

func main() {
	var (
		webserver bool
		method    string
		url       string
	)

	flag.BoolVar(&webserver, "webserver", false, "run as webserver")
	flag.StringVar(&method, "X", "", "method")
	flag.Parse()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	if webserver {
		log.Fatal(http.ListenAndServe(":8080", router))
		return
	}

	if _, ok := supportedMethods[method]; !ok {
		log.Fatal(fmt.Sprintf("Invalid method: %s", method))
	}

	args := flag.Args()
	if len(args) == 0 {
		url = "/"
	} else {
		url = args[0]
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot create request: %s", err.Error()))
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	fmt.Print(resp.Body.String())
}
