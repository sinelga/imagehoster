package handlers

import (
"net/http"
	"github.com/zenazn/goji/web"
	"log/syslog"
	"startones"
	"sync"
)

var startOnce sync.Once
var golog syslog.Writer


func MhandleAll(c web.C, w http.ResponseWriter, r *http.Request) {
		
		startOnce.Do(func() {
		golog ,_ = startones.Start()

	})	
		
			
}