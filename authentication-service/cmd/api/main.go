package main

import (
	"authentication/config"
	database "authentication/connection"
	"authentication/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Auth service is on")
	conn := database.ConnectToDB()
	if conn == nil {
		log.Fatal("Could not connect to database")
	}
	// add migration
	app := config.NewConfig(conn, "80")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", app.Webport),
		Handler: routes.GetMux(app.DB),
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
