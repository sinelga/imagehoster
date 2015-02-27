package main

import (
	"database/sql"
	"flag"
	//	"domains"
	_ "github.com/go-sql-driver/mysql"
	//	"log/syslog"
	"startones"
	//    "fmt"
	//    "domains"
	"contents_feeder_redis/find_all_img"
)

var siteFlag = flag.String("site", "", "must be test.com www.test.com")

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	golog, config := startones.Start()

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()
	

	//    site := *siteFlag
	
//	names := FindAll(golog, *db)
		
	var characters find_all_img.Characters
	//
	characters.Find_all_img("/home/juno/git/imagehoster/upload/img")
	
	
	
	

}
