package html_server

import (
    "time"
    "fmt"
    "net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v: Recieved request for HelloWorld\n", time.Now())
	w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "Hello World!")
}
