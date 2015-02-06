package getAll

import (
	
	"github.com/zenazn/goji/web"
	"net/http"
	"log/syslog"
	"log"

)

func GetAll(c web.C, w http.ResponseWriter, r *http.Request) {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	golog.Info("Start GetAll")
	
	
}

