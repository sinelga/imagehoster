package find_names

import (
	"database/sql"
	"domains"
	_ "github.com/go-sql-driver/mysql"
	"log/syslog"
)

func FindAll(golog syslog.Writer, config domains.Config) []string {

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select name from girl_names")

	var names []string

	for rows.Next() {

		var name string

		if err := rows.Scan(&name); err != nil {

			golog.Err(err.Error())

		} else {
			names = append(names, name)

		}

	}

	return names
}
