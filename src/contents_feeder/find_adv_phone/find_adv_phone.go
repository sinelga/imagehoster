package find_adv_phone

import (
	"database/sql"
	"domains"
	_ "github.com/go-sql-driver/mysql"
	"log/syslog"
)

func FindAll(golog syslog.Writer, config domains.Config) []int {

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select id from adv_phones")

	var adv_phone_id []int

	for rows.Next() {

		var id int
		if err := rows.Scan(&id); err != nil {

			golog.Err(err.Error())
			
		} else {
			adv_phone_id = append(adv_phone_id, id)

		}

	}
	
	return adv_phone_id

}
