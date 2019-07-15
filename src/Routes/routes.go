package Routes

import "github.com/gin-gonic/gin"

func Router() {

	router := gin.Default()

	v1 := router.Group("/CRUD")
	{
		v1.POST("/", createUrl)
		v1.GET("/", fetchAllUrl)
		v1.GET("/:id", fetchUrlLog)

		// 	v1.PUT("/:id", updateUrl)
		// 	v1.DELETE("/:id", deleteUrl)

	}
	router.Run()
}