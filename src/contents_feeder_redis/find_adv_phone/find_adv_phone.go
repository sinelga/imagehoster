package find_adv_phone

import (
	"database/sql"
	"log/syslog"
)

func FindAll(golog syslog.Writer, db sql.DB) []string {

	var adv_phone []string

	if rows, err := db.Query("select phone from adv_phones"); err != nil {
		golog.Crit(err.Error())

	} else {

		for rows.Next() {

			var phone string
			if err := rows.Scan(&phone); err != nil {

				golog.Err(err.Error())

			} else {
				adv_phone = append(adv_phone, phone)

			}

		}

	}

	return adv_phone

}
