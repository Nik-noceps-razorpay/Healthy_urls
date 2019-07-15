package Models

import "github.com/jinzhu/gorm"







type UrlHits struct {
	gorm.Model
	Hit_number int
	Status int
	UrlId uint
}