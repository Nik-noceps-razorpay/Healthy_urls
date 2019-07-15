package Routes

import (
	"Controllers"
	"github.com/gin-gonic/gin"
)


func Router() {

	router := gin.Default()

	v1 := router.Group("/CRUD")
	{
		v1.POST("/", Controllers.CreateUrl)
		v1.GET("/", Controllers.FetchAllUrl)
		v1.GET("/:id", Controllers.FetchUrlLog )

		// 	v1.PUT("/:id", updateUrl)
		// 	v1.DELETE("/:id", deleteUrl)

	}
	router.Run()
}




