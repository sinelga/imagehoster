package find_names

import (
	"database/sql"
//	"domains"
//	_ "github.com/go-sql-driver/mysql"
	"log/syslog"
)

func FindAll(golog syslog.Writer, db sql.DB) []string {


var names []string

 if rows, err := db.Query("select name from girl_names"); err != nil  {
 	
 	golog.Crit(err.Error())
 	
 } else {

	for rows.Next() {

		var name string

		if err := rows.Scan(&name); err != nil {

			golog.Err(err.Error())

		} else {
			names = append(names, name)

		}

	}
	
	}

	return names
}
