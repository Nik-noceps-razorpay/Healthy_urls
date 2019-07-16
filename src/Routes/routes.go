package Routes

import (
	"Controllers"
	"github.com/gin-gonic/gin"
)


func Router() {

	router := gin.Default()

	v1 := router.Group("/healthyurls")
	{
		v1.POST("/CRUD", Controllers.CreateUrl)
		v1.GET("/CRUD", Controllers.FetchAllUrl)
		v1.GET("/CRUD/:id", Controllers.FetchUrlLog)
		v1.GET("/readfile", Controllers.ReadUrl)

		// 	v1.PUT("/:id", updateUrl)
		// 	v1.DELETE("/:id", deleteUrl)

	}
	router.Run()
}




