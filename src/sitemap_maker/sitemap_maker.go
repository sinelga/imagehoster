package main

import (
	"database/sql"
	"domains"
	"flag"
	"fmt"
	"net/url"
	"time"
	//	"fmt"
	"encoding/xml"
	_ "github.com/go-sql-driver/mysql"
	//	"log/syslog"

	//    "domains"

	//	"log/syslog"

	"startones"
	"strconv"

	"sitemap_maker/getLinks"
)

//const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
//var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var domainFlag = flag.String("domain", "", "must be test.com www.test.com")
var limitFlag = flag.Int("limit", 0, "if not will be 0")

func main() {
	flag.Parse() // Scan the arguments list

	//	if *versionFlag {
	//		fmt.Println("Version:", APP_VERSION)
	//	}

	domain := *domainFlag
	limit := *limitFlag

	golog, config := startones.Start()

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	characters := getLinks.GetAllLinks(golog, *db, limit)

	var Url *url.URL

	docList := new(domains.Pages)
	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for _, character := range characters {

		Url, err = url.Parse("http://"+domain)
		if err != nil {
			panic("boom")
		}

		Url.Path += "/" + strconv.Itoa(character.Id) + "/" + character.Moto

		doc := new(domains.Page)
		doc.Loc = Url.String()
		doc.Lastmod = time.Now().Local().Format(time.RFC3339)
		doc.Changefreq="daily" 

		docList.Pages = append(docList.Pages, doc)

	}

	resultXml, err := xml.MarshalIndent(docList, "", "  ")
	if err != nil {

		golog.Crit(err.Error())
	}

	fmt.Println(string(resultXml))

}
