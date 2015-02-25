package handlers

import (
//	"fmt"
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
	"strconv"
	
)

func ImageShow(c web.C, w http.ResponseWriter, r *http.Request) {
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	golog, config = startones.Start()
	width := 100;
	height := 100;
	

	id := c.URLParams["id"]
	imgfile := c.URLParams["imgfile"]
	mime := c.URLParams["mime"]
	
	if c.URLParams["width"] !="" {
		
		widthstr :=c.URLParams["width"]
		width,_ = strconv.Atoi(widthstr)
			
	}
	
	if c.URLParams["height"] !="" {
		
		heightstr :=c.URLParams["height"]
		height,_ = strconv.Atoi(heightstr)
			
	}	

	w.Header().Set("Content-Type", "image/"+mime)

	filestr := config.Store.StoreDir + id + "/original/" + imgfile

//	fmt.Println(filestr)

	file, err := imaging.Open(filestr)
//	defer file.Close()
	if err != nil {
		//			log.Println(err.Error())
		golog.Err(err.Error()+" "+filestr)
		return
	}

	m := imaging.Thumbnail(file,width ,height , imaging.CatmullRom)
	
	jpeg.Encode(w, m, nil)

}
