package main

import (
	"fmt"
	"github.com/vanilla/go-mux-crud/api/common"
	log "github.com/sirupsen/logrus"
	"github.com/vanilla/go-mux-crud/api/router"
	"net/http"
)

func init() {
	var err error

	err = common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Printf("Listening on port %s\n", common.Config.Port)

	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", common.Config.Port), r))
}