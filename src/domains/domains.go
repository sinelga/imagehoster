package domains

import (
"time"
)

type Config struct {
	Database struct {
		ConStr string
	}
}


type Character struct {
	
	Id int
	Name string
	Age int
	Moto string
	Description string
	Region_id int
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