package main

import (
//	"net/http"
	"log"
	"log/syslog"
	"github.com/rs/cors"
	"github.com/zenazn/goji"
	"handlers"
	
)


func main() {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	goji.Use(c.Handler)

	goji.Get("/api", handlers.MhandleAll)
	
	goji.Get("/img/:id/:imgfile/:img/:mime/:width/:height", handlers.ImageShow)
	
	
	goji.Options("/upload",handlers.MakeUpload)	
	goji.Post("/upload",handlers.MakeUpload) 
	

	goji.Serve()
}
