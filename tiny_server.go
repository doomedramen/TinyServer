package main
import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "port to run on")
	flag.Parse()
	portString := fmt.Sprintf(":%d", *port)
	panic(http.ListenAndServe(portString, http.FileServer(http.Dir("./"))))
}
