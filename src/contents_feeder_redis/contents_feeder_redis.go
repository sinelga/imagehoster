package main

import (
	"database/sql"
	"flag"
	//	"domains"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	//	"log/syslog"
	"fmt"
	"startones"
	//    "domains"
	"contents_feeder_redis/find_adv_phone"
	"contents_feeder_redis/find_all_img"
	"contents_feeder_redis/find_names"
	"contents_feeder_redis/find_regions"
)

var siteFlag = flag.String("site", "", "must be test.com www.test.com")
var deltahoursFlag = flag.Int("deltahours", 0, " > 0")


func main() {
	flag.Parse() // Scan the arguments list

	golog, config := startones.Start()

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	redisprotocol := config.Redis.Prot
	redishost := config.Redis.Host

	c, err := redis.Dial(redisprotocol, redishost)
	if err != nil {

		golog.Crit(err.Error())

	}
	defer c.Close()
	
	site := *siteFlag
	deltahours :=*deltahoursFlag
	
	if site !="" && deltahours > 0{	 

	names := find_names.FindAll(golog, *db)
	phones := find_adv_phone.FindAll(golog, *db)
	regions := find_regions.FindAll(golog, *db)

	var ch find_all_img.Characters
	//
	ch.Find_all_img("/home/juno/git/imagehoster/upload/img")

	ch.Add_name_phone_region(golog, names, phones, regions)

	ch.Find_age(golog, *db)

	ch.Find_free_paragraph(golog, c, "fi_FI", "porno")
	
	ch.Create_local_charters(golog,site,deltahours)

//	for i, _ := range ch.CharactersRedis {
//
//		fmt.Println(ch.CharactersRedis[i].Name, ch.CharactersRedis[i].Phone, ch.CharactersRedis[i].City, ch.CharactersRedis[i].Age,ch.CharactersRedis[i].Moto,ch.CharactersRedis[i].Description)
//		
//	}

	} else {
		
		fmt.Println("Need paramater --site=wwww.test.com --deltahours=240")
		
		
	}

}
