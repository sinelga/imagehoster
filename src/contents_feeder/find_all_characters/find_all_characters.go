package find_all_characters

import (
	"domains"
	"log/syslog"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func FindAll(golog syslog.Writer,config domains.Config) []domains.Character{

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select Id,Name,Age,Moto,Description,Region_id,Adv_phone_id,Img_orient,Topic,Sex, Created_at,Updated_at,Img_file_name,Img_content_type,Img_file_size,Img_updated_at from characters where (Moto='' or Description ='') and  Topic='sex' and sex='female'")
	if err != nil {
		golog.Err(err.Error())
	}
	defer rows.Close()

	//	fmt.Println("N rows", len(rows))
	var characters []domains.Character

	for rows.Next() {

		var ch domains.Character

		err := rows.Scan(&ch.Id, &ch.Name, &ch.Age, &ch.Moto, &ch.Description, &ch.Region_id, &ch.Adv_phone_id, &ch.Img_orient, &ch.Topic, &ch.Sex, &ch.Created_at, &ch.Updated_at, &ch.Img_file_name, &ch.Img_content_type, &ch.Img_file_size, &ch.Img_updated_at)
		if err != nil {
			golog.Err(err.Error())
		}

		characters = append(characters, ch)

	}
	
	
	return characters 

}
