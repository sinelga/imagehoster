package find_adv_phone

import (
	"database/sql"
	"fmt"
	"startones"
	"testing"

	//	"domains"
	_ "github.com/go-sql-driver/mysql"
)

func TestXYZ(t *testing.T) {

	golog, config := startones.Start()

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	names := FindAll(golog, *db)

	for _, name := range names {

		fmt.Println(name)

	}

}
