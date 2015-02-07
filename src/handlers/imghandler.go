package handlers

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"github.com/nfnt/resize"
)

func ImageShow(c web.C, w http.ResponseWriter, r *http.Request) {

	id := c.URLParams["id"]
	imgfile := c.URLParams["imgfile"]
	mime := c.URLParams["mime"]

	w.Header().Set("Content-Type", "image/"+mime)

	filestr := "upload/img/" + id + "/original/" + imgfile

	fmt.Println(filestr)

	file, err := os.Open(filestr)
	defer file.Close()
	if err != nil {
		//			log.Println(err.Error())
		http.Error(w, "not found", 404)
		return
	}

	var img image.Image

	if mime == "jpeg" {
		img, err = jpeg.Decode(file)
	}
	if mime == "png" {
		img, err = png.Decode(file)
	}

	if err != nil {
		//			log.Println(err.Error())
		http.Error(w, "not found", 404)
		return
	}

//	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
//	m := resize.Resize(uint(200), uint(200), img, resize.Lanczos3)
	m := resize.Thumbnail(uint(200), uint(0),img, resize.Lanczos3)

	jpeg.Encode(w, m, nil)

}
