package main

import (
	"fmt"
	"log"
	"net/http"
	"rentx/src/framework/router"
	"rentx/src/utils/config"
)

func main() {
	config.LoadEnvs()

	fmt.Printf("Listening at port %d!\n", config.Port)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
