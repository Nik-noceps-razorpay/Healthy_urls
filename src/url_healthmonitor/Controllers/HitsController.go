package Controllers

import (
	"DbConn"
	"Models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func FetchUrlLog(c *gin.Context) {

	id := c.Param("id")

	fmt.Println(id)
	var u []Models.Url

	DbConn.Db.Model(&Models.Url{}).Where("id = ?", id).First(&u)

	var hist []Models.UrlHits

	DbConn.Db.Model(&Models.UrlHits{}).Where("url_id = ?", u[0].ID).Order("hit_number").Find(&hist)

	for i := 0; i < len(hist) ; i++ {

		fmt.Println(hist[i])

	}

}
