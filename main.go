package main

// package for building web services
import "net/http"

func main() {
	// bind this to every IP on machine at port 9090
	// second parameter is a handler, ignore this now
	http.ListenAndServe(":9090", nil)
}
