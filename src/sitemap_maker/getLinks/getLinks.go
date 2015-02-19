package getLinks

import (
	"database/sql"
	"domains"
	_ "github.com/go-sql-driver/mysql"
	"log/syslog"
)

func GetAllLinks(golog syslog.Writer, db sql.DB) []domains.Character {

	sqlstr := "select Id,Moto from characters where topic='sex' and sex='female' order by Created_at desc limit 50"

	rows, err := db.Query(sqlstr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer rows.Close()

	var characters []domains.Character

	for rows.Next() {

		var ch domains.Character

		err := rows.Scan(&ch.Id, &ch.Moto)
		if err != nil {
			golog.Err(err.Error())
		}

		characters = append(characters, ch)

	}

	return characters
}