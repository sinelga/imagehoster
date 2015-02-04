package main
import "encoding/json"
import "database/sql"
import "path/filepath"
import "fmt"
import "github.com/go-martini/martini"
import "code.google.com/p/gcfg"
import "github.com/martini-contrib/render"
import "github.com/nfnt/resize"
import _ "github.com/go-sql-driver/mysql"
import "io"
import "io/ioutil"
import "image/jpeg"
import "image/png"
import "image"
import "log"
import "net/http"
import "os"
import "flag"
import "strconv"
import "mime"



type DbImage struct {
	images_id int64
	path string
	shorturl string
	mime string
}


func getImage(shorturl string) (*DbImage, error)  {
		db,err := sql.Open("mysql", config.Database.ConStr)
		defer db.Close()
		
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		row := db.QueryRow("select * from images where shorturl = ?", shorturl);
		image := new(DbImage)
		err = row.Scan(&image.images_id, &image.path, &image.shorturl, &image.mime)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return image, nil

}
type Config struct {
	Database struct {
		ConStr string
	}
}
var config Config
func main() {

	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		log.Println("cannot read configuration file config.ini", err.Error())
		return
	}

	log.Println("parsed db ConStr" , config.Database.ConStr)
	store := flag.String("store", "", "image store path");
	flag.Parse();
	if *store == "" {
		fmt.Println("-store was not given, usage <program> -store <path>")
		return 
	}
	fmt.Println("Store: ", *store)
	server := martini.Classic()
	server.Use(render.Renderer())
	server.Use(martini.Static("assets"))
	server.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "deniz")
	})

	server.Get("/serv/:shorturl/:w/:h", func(params martini.Params, res http.ResponseWriter) {
		dbimage, err := getImage(params["shorturl"])
		res.Header().Set("Content-Type", dbimage.mime)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "not found", 404)
			return
		}
		file,err := os.Open(dbimage.path)
		defer file.Close()
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "not found", 404)
			return
		}

	    var img image.Image
		if dbimage.mime == "image/jpeg" {
			img,err = jpeg.Decode(file)
		}
		if dbimage.mime == "image/png" {
			img,err = png.Decode(file)
		}

		if err != nil {
			log.Println(err.Error())
			http.Error(res, "not found", 404)
			return
		}
		width,err := strconv.ParseUint(params["w"], 10, 0)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "width parse error", 500)
			return
		}
		height,err := strconv.ParseUint(params["h"], 10, 0)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "height parse error", 500)
			return
		}
		m := resize.Resize(uint(width), uint(height),img, resize.Lanczos3)
//		m := resize.Thumbnail(uint(width), uint(height),img, resize.Lanczos3)
		
		jpeg.Encode(res,m,nil)
	})
	
	server.Get("/serv/:shorturl", func(params martini.Params, res http.ResponseWriter) {
		image, err := getImage(params["shorturl"])

		res.Header().Set("Content-Type", image.mime)

		buf,err := ioutil.ReadFile(image.path) 
		if err != nil {
			log.Println(err.Error())
			http.Error(res, "not found", 404)
			return
		}
		res.Write(buf)
	})
	
	server.Post("/upload", func(r *http.Request, res http.ResponseWriter) (int,string) {
		log.Println("parsing request")
	
		err := r.ParseMultipartForm(10000000)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		files := r.MultipartForm.File["file"]
		results := make(map[string]string)
		for i,_ := range files {
			file,err := files[i].Open()
			defer file.Close()

			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}
			
			log.Println("creating file")
			dst, err := os.Create(*store + "/" + files[i].Filename)
			defer dst.Close()

			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			if _,err := io.Copy(dst,file); err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			log.Println("db str", config.Database.ConStr)
			db,err := sql.Open("mysql", config.Database.ConStr)
			defer db.Close()

			ext := filepath.Ext(files[i].Filename)
			mime := mime.TypeByExtension(ext)
			log.Println("type", mime)
			res , err := db.Exec("insert into images (images_id,path,mime) values (null,?,?)", *store+ "/"+files[i].Filename,mime)
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			id, err := res.LastInsertId()
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			shorturl := getShortUrl(id)
			results[files[i].Filename] = shorturl
			if _, err := db.Exec("update images set shorturl = ? where images_id = ?", shorturl, id); err != nil {
				return http.StatusInternalServerError, err.Error()
			}
			
		}
		jsn,err := json.Marshal(results)
		res.Header().Set("Content-Type", "application/json")
		return 200, string(jsn)
	})
	server.Run()
}

func getShortUrl(id int64) string {
	return fmt.Sprintf("%x", id)
}
