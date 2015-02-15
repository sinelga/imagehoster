package startones

import (
	//	"github.com/garyburd/redigo/redis"
	"code.google.com/p/gcfg"
//	"io/ioutil"
	"log"
	"log/syslog"
	"domains"

)



var config domains.Config

//func Start(golog syslog.Writer) ([]string,map[string]struct{}) {
func Start() (syslog.Writer, domains.Config) {

	//	sitestoblock := make(map[string]struct{})

	golog, err := syslog.New(syslog.LOG_ERR, "golog")	

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

//	golog.Info("StartOnes")
	
	err = gcfg.ReadFileInto(&config, "/home/juno/git/imagehoster/config.ini")
	if err != nil {
		
		golog.Crit("cannot read configuration file config.ini" + err.Error())
//		return nil,nil
	}
	
	golog.Info(config.Database.ConStr)



	return *golog, config

}
