package handlers

import (
	"domains"
	"encoding/json"
	"github.com/zenazn/goji/web"
	"handlers/getAll"
	"log/syslog"
	"net/http"
	"startones"
	"sync"
)

var startOnce sync.Once
var golog syslog.Writer

var config domains.Config

func MhandleAll(c web.C, w http.ResponseWriter, r *http.Request) {

	startOnce.Do(func() {
		golog, config = startones.Start()

	})

	golog.Info(c.URLParams["id"])
	golog.Info(r.Method)


	characters := getAll.GetAll(golog, config)

	bytes, e := json.Marshal(characters)
	if e != nil {

		golog.Err(e.Error())

	}
	w.Write(bytes)

}
