package find_regions

import (
	"database/sql"
//	"domains"
//	_ "github.com/go-sql-driver/mysql"
	"log/syslog"
)

func FindAll(golog syslog.Writer, db sql.DB) []string {

var citys []string

if 	rows, err := db.Query("select city from regions"); err !=nil {
	golog.Crit(err.Error())
	
	
} else {

	

	for rows.Next() {

		var city string
		if err := rows.Scan(&city); err != nil {

			golog.Err(err.Error())

		} else {
			citys = append(citys, city)

		}

	}
	
	}

	return citys

}
