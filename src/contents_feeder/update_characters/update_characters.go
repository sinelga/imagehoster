package update_characters

import (
	"database/sql"
	"domains"
	_ "github.com/go-sql-driver/mysql"
	"log/syslog"
	//	"fmt"
)

func Update(golog syslog.Writer, config domains.Config, characters []domains.Character) {

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		golog.Err(err.Error())
	}

	//	fmt.Println("Character ",len(characters))

	for _, character := range characters {

		golog.Info(character.Name)

		if stmt, err := tx.Prepare("update characters set name=?,moto=?,description=?,region_id=?, adv_phone_id=? where id=?"); err != nil {

			golog.Err(err.Error())

		} else {
			defer stmt.Close()

			if _, err = stmt.Exec(character.Name, character.Moto, character.Description, character.Region_id, character.Adv_phone_id, character.Id); err != nil {

				golog.Err(err.Error())

			}

		}

		//		stmt, err := db.Prepare("update characters set name=?,moto=?,descriptios=?, where uid=?")

	}

	//		stmt.Close()
	tx.Commit()

}
