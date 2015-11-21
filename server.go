package main

import (
	"flag"
	"github.com/MusculoTherm/musclesapi/controllers"
	"github.com/MusculoTherm/musclesapi/models"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var addr = flag.String("addr", ":8080", "HTTP Service Address")

func main() {
	flag.Parse()
	err := models.Setup()
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(*addr, controllers.CreateRouter())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
