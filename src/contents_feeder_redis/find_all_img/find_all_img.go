package find_all_img

import (
	"domains"
	//	"fmt"
	"database/sql"
	"log/syslog"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Characters struct {
	CharactersRedis []domains.CharacterRedis
}

var imgfiles [][]string

func walkpath(pathstr string, f os.FileInfo, err error) error {

	if !f.IsDir() {
	
		dir, file := path.Split(pathstr)
		dirscplit := strings.Split(dir, "/")
		id := dirscplit[len(dirscplit)-3]
		imgfile := []string{id, file}
		imgfiles = append(imgfiles, imgfile)

	}
	return nil
}

func (characters *Characters) Find_all_img(dir string) {

	filepath.Walk(dir, walkpath)

	for _, imgfile := range imgfiles {

		var character domains.CharacterRedis

		character.Img_file_name = imgfile[1]
		idint, _ := strconv.Atoi(imgfile[0])
		character.Id = idint

		characters.CharactersRedis = append(characters.CharactersRedis, character)

	}

}

func (characters *Characters) Add_name_phone_region(golog syslog.Writer, names []string, phones []string, regions []string) {

	names_quant := len(names)
	phones_quant := len(phones)
	regions_quant := len(regions)

	rand.Seed(time.Now().UTC().UnixNano())

	now := time.Now().Local()

	for i, _ := range characters.CharactersRedis {

		int_name := rand.Intn(names_quant)
		int_phone := rand.Intn(phones_quant)
		int_ragions := rand.Intn(regions_quant)

		characters.CharactersRedis[i].Name = names[int_name]
		characters.CharactersRedis[i].Phone = phones[int_phone]
		characters.CharactersRedis[i].City = regions[int_ragions]
		characters.CharactersRedis[i].Created_at = now

	}

}

func (characters *Characters) Find_age(golog syslog.Writer, db sql.DB) {

	for i, _ := range characters.CharactersRedis {

		sqlstr := "select age from characters where id=" + strconv.Itoa(characters.CharactersRedis[i].Id)

		if rows, err := db.Query(sqlstr); err != nil {
			golog.Crit(err.Error())

		} else {

			for rows.Next() {

				var age int
				if err := rows.Scan(&age); err != nil {

					golog.Err(err.Error())

				} else {
					
					characters.CharactersRedis[i].Age=age

				}

			}

		}

	}

}
