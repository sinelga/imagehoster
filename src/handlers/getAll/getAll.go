package getAll

import (
	"database/sql"
	"domains"
//	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/syslog"
	"time"
)

func GetAll(golog syslog.Writer, config domains.Config) []domains.Character {

	golog.Info("Start GetAll" + time.Now().String())

	golog.Info(config.Database.ConStr)

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	sqlstr := "select ch.Id,Name,Age,Moto,ph.Phone,Description,Region_id,City,Adv_phone_id,Img_orient,Topic,Sex,ch.Created_at,ch.Updated_at,Img_file_name,Img_content_type,Img_file_size,Img_updated_at from characters as ch,adv_phones as ph,regions as re where re.id=ch.region_id and ph.id=ch.adv_phone_id and topic='sex' and sex='female' order by  RAND() limit 10"

	rows, err := db.Query(sqlstr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer rows.Close()

	//	fmt.Println("N rows", len(rows))
	var characters []domains.Character

	for rows.Next() {

		var ch domains.Character

		err := rows.Scan(&ch.Id,&ch.Name,&ch.Age,&ch.Moto,&ch.Phone,&ch.Description,&ch.Region_id,&ch.City,&ch.Adv_phone_id,&ch.Img_orient,&ch.Topic,&ch.Sex, &ch.Created_at,&ch.Updated_at,&ch.Img_file_name,&ch.Img_content_type,&ch.Img_file_size,&ch.Img_updated_at)
		if err != nil {
			golog.Err(err.Error())
		}

		characters = append(characters, ch)

	}
	

	return characters 

}
