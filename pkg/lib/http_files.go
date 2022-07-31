package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type MySrv struct{}

func (s *MySrv) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	switch uri := r.RequestURI; uri {
	default:
	case "/":
	case "/files":
	}
}

func HTTPFiles() {
	var uIn string
	srvErr := make(chan error)
	srv := &MySrv{}

	go func() {
		fmt.Println("Starting the server on port 8000")
		if err := http.ListenAndServe(":8000", srv); err != nil {
			srvErr <- err
		}
	}()
	fmt.Println("Type 'q' or 'quit' to stop the app.")
	fmt.Scanln(&uIn)
	if uIn == "q" || uIn == "quit" {
		fmt.Println("Gracefully shutting down.")
		os.Exit(0)
	}
	log.Fatal(<-srvErr)
}
