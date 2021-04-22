package api

import (
	"blogapi/api/router"
	"blogapi/config"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	fmt.Printf("server running.. at port %d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
