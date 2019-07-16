package Controllers

import (
	"Models"
	"encoding/json"
	"io/ioutil"

	"DbConn"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

var Db *gorm.DB

func CreateUrl(c *gin.Context) {
	var x []Models.Url
	c.Bind(&x)
	fmt.Println(x)
	fmt.Println("\n\n")
	for i := 0; i < len(x); i++ {
		go Put(x[i])
	}
}

func ReadUrl(c *gin.Context){

	var x []Models.Url
	p, _ := ioutil.ReadFile(c.Query("path"))
	json.Unmarshal(p, &x)

	//fmt.Println("route ho rha hai", c.Query("path"))
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
		go Put(x[i])
	}

}

func Put(x Models.Url) {
	fmt.Println("here")
	var count int
	var u Models.Url

	DbConn.Db.Model(&Models.Url{}).Where("url_name = ?",x.UrlName).Count(&count)

	if count !=0 {

		DbConn.Db.Where("url_name= ?",x.UrlName).First(&u)

		u.Frequency = x.Frequency
		u.Crawl_timeout = x.Crawl_timeout
		u.Failure_threshold = x.Failure_threshold

		DbConn.Db.Save(&u)

		fmt.Println("Url", u.UrlName, "has been updated")
	} else {

		DbConn.Db.Save(&x)

		fmt.Println("inserting data into table Url ")

	}

}


func FetchAllUrl(c *gin.Context) {
	var urls []Models.Url

	DbConn.Db.Find(&urls)

	if len(urls) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No urls found!, kindly insert some urls."})
		return
	}
	for i := 0 ; i < len(urls) ; i++ {
		fmt.Println("urls are :", urls[i] )

	}

}