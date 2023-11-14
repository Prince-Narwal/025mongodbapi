package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prince/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("MongoDB API")
	fmt.Println("server is getting starte...")
	log.Fatal(http.ListenAndServe(":4000", r))

}
