package main

import (
	simplecli "github.com/sifatulrabbi/hello-go/simple-cli"
	// "fmt"
	// "log"
	// "net/http"
	// "github.com/gorilla/mux"
	// simplecrudserver "github.com/sifatulrabbi/hello-go/simple_crud_server"
)

func main() {
	simplecli.SimpleCLI()
	// go helloServer()
	// simplecrudserver.SimpleCrudServer()
}

// func helloServer() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/hello", handleHello)
// 	fmt.Printf("Starting server on port 8001")
// 	log.Fatal(http.ListenAndServe(":8001", router))
// }

// func handleHello(wr http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(wr, "Hello, world!")
// 	wr.WriteHeader(http.StatusOK)
// }
