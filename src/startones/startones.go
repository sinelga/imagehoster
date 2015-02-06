package startones

import (
	//	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"log/syslog"
	//	"os"
	"strings"
	//	"fmt"
)

//func Start(golog syslog.Writer) ([]string,map[string]struct{}) {
func Start() (syslog.Writer, []string) {

	//	sitestoblock := make(map[string]struct{})

	golog, err := syslog.New(syslog.LOG_ERR, "golog")	

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	golog.Info("StartOnes")

	content, err := ioutil.ReadFile("/home/juno/git/imagehoster/config.txt")
	if err != nil {
		//Do something
		golog.Err(err.Error())
	}
	parameters := strings.Split(string(content), ",")
	cleanparameters := []string{strings.TrimSpace(parameters[0]), strings.TrimSpace(parameters[1]), strings.TrimSpace(parameters[2])}

	return *golog, cleanparameters

}
