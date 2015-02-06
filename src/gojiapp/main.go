package main

import (
	"net/http"
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

	goji.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	goji.Get("/api", handlers.MhandleAll)

	goji.Serve()
}
