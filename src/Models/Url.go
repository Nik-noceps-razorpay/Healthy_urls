package Models

import "github.com/jinzhu/gorm"







type Url struct {
	gorm.Model
	UrlName           string `gorm:"unique;not null" json:"url_name"`
	Crawl_timeout     int    `json:"crawl_timeout`
	Frequency         int    `json:frequency`
	Failure_threshold int    `json:failure_threshold`
	Health            int    `gorm:"default:2"`
}