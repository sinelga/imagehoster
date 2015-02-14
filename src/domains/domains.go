package domains

import (
"time"
)

type Config struct {
	Database struct {
		ConStr string
	}
	Store struct {
		StoreDir string
	}
	Redis struct {
		Prot string
		Host string
		
	}
}


type Character struct {
	
	Id int
	Name string
	Age int
	Moto string
	Description string
	City string
	Region_id int
	Phone string
	Adv_phone_id int
	Img_orient string	
	Topic string
	Sex string
	Created_at time.Time
	Updated_at time.Time
	Img_file_name string
	Img_content_type string
	Img_file_size int
	Img_updated_at time.Time
}

type Paragraph struct {
	Ptitle     string
	Pphrase    string
	Plocallink string
	Phost      string
	Sentences  []string
	Pushsite   string
}