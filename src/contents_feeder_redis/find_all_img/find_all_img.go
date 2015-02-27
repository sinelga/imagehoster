package find_all_img

import (
	"domains"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Characters struct {
	CharactersRedis []domains.CharacterRedis
}

var imgfiles [][]string





func walkpath(pathstr string, f os.FileInfo, err error) error {

	if !f.IsDir() {
		//		fmt.Printf("%s with %d bytes\n", pathstr, f.Size())
		dir, file := path.Split(pathstr)
		fmt.Println(dir)
		dirscplit := strings.Split(dir, "/")
		//		fmt.Println(dirscplit[len(dirscplit)-3])
		id := dirscplit[len(dirscplit)-3]
		fmt.Println(id)
		fmt.Println(file)

		imgfile := []string{id, file}

		imgfiles = append(imgfiles, imgfile)

	}
	return nil
}

func (characters *Characters) Find_all_img(dir string) {

	filepath.Walk(dir, walkpath)

	for _, imgfile := range imgfiles {

		//		fmt.Println(imgfile[0])

		var character domains.CharacterRedis

		character.Img_file_name = imgfile[1]

		idint, _ := strconv.Atoi(imgfile[0])
		character.Id = idint

		characters.CharactersRedis = append(characters.CharactersRedis, character)

		//		characters =append(characters,character)

	}

}

func (characters *Characters) Add_name_location() {

	for _, ch := range characters.CharactersRedis {
		fmt.Println(ch.Img_file_name)

	}

}
