package handlers

import (
	"fmt"
	"github.com/zenazn/goji/web"
//	"image"
	"image/jpeg"
//	"image/png"
	"net/http"
//	"os"
//	"github.com/nfnt/resize"
	"github.com/disintegration/imaging"
	"runtime"
	"startones"
	
)

func ImageShow(c web.C, w http.ResponseWriter, r *http.Request) {
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	golog, config = startones.Start()

	id := c.URLParams["id"]
	imgfile := c.URLParams["imgfile"]
	mime := c.URLParams["mime"]

	w.Header().Set("Content-Type", "image/"+mime)

	filestr := config.Store.StoreDir + id + "/original/" + imgfile

	fmt.Println(filestr)

	file, err := imaging.Open(filestr)
//	defer file.Close()
	if err != nil {
		//			log.Println(err.Error())
		golog.Err(err.Error()+" "+filestr)
		return
	}

	m := imaging.Thumbnail(file, 100, 100, imaging.CatmullRom)
	
	jpeg.Encode(w, m, nil)

}
