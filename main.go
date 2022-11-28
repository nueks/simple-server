package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// todo
// return envvariable
// return count
func handler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, fmt.Sprintf("hostname: %v\n", host))
	fmt.Println(fmt.Sprintf("timestamp=%s method=%s route=%s protocol=%s status_code=%d",
		time.Now().Format("2006-01-02T15:04:05"),
		r.Method,
		r.URL,
		r.Proto,
		200))
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "80"
	}
	log.Printf("Listening on :%v...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
