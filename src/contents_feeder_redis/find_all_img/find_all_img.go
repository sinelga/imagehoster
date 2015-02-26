package find_all_img

import (
	"fmt"
	"os"
	"path/filepath"
	"path"
	"strings"
)

func walkpath(pathstr string, f os.FileInfo, err error) error {

	if !f.IsDir() {
//		fmt.Printf("%s with %d bytes\n", pathstr, f.Size())
		dir, file := path.Split(pathstr)
		fmt.Println(dir)
		dirscplit :=strings.Split(dir,"/")
//		fmt.Println(dirscplit[len(dirscplit)-3])
		id := dirscplit[len(dirscplit)-3]
		fmt.Println(id) 
		 
		
		fmt.Println(file)
		

	}
	return nil
}

func Find_all_img(dir string) {

	filepath.Walk(dir, walkpath)

}
