package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)


func main() {
	var (
		listenFlag = flag.String("listen", EnvOrDefault("SIMPLE_WEBSERVER_LISTEN", ":8082"), "Address + Port to listen on. Format ip:port. Environment variable: SIMPLE_WEBSERVER_LISTEN")
	)
	flag.Parse()

	// Define HTTP endpoints
	s := http.NewServeMux()
	s.HandleFunc("/payload", PayloadHandler)

	// Bootstrap logger
	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Printf("Starting webserver and listen on %s", *listenFlag)

	// Start HTTP Server with request logging
	loggingHandler := handlers.LoggingHandler(os.Stdout, s)
	log.Fatal(http.ListenAndServe(*listenFlag, loggingHandler))
}


// PayloadHandler handles request to the "/payload" endpoint.
// It is a debug route to dump the complete request incl. method, header and body.
func PayloadHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	log.Printf("Method: %s\n", req.Method)
	fmt.Fprintf(resp, "Method: %s\n", req.Method)

	log.Printf("Src IP: %s\n", req.RemoteAddr)
	fmt.Fprintf(resp, "src IP: %s\n", req.RemoteAddr)

	log.Printf("Date: %s\n", time.Now())
	fmt.Fprintf(resp, "Date: %s\n", time.Now())
	host,_ := os.Hostname()
	log.Printf("Hostname: %s\n", host)
	fmt.Fprintf(resp, "Hostname: %s\n",host)
}

// EnvOrDefault will read env from the environment.
// If the environment variable is not set in the environment
// fallback will be returned.
// This function can be used as a value for flag.String to enable
// env var support for your binary flags.
func EnvOrDefault(env, fallback string) string {
	value := fallback
	if v := os.Getenv(env); v != "" {
		value = v
	}

	return value
}
