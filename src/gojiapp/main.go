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
	goji.Get("/api/:id", handlers.MhandleAll)	
	goji.Get("/img/:id/:imgfile/:img/:mime/:width/:height", handlers.ImageShow)
	goji.Get("/chat/:uuid/:phone/:ask",handlers.GetChatAnswer)
	goji.Get("/uploadapi", handlers.ShowAll)
//	goji.Get("/fullimage/", http.FileServer(http.Dir("upload/img/")))
	goji.Get("/fullimage/:id/original/:imgfile", handlers.ImageFullShow)
//	goji.Get("/fullimage/:id/original/:imgfile",http.FileServer(http.Dir("upload/img"))) 
//	goji.Get("/fullimage",handlers.ImageFullShow)
	
	
	goji.Options("/upload",handlers.MakeUpload)	
	goji.Post("/upload",handlers.MakeUpload) 
	

	goji.Serve()
}
