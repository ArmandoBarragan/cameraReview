package main

import (
	"cameraReview/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Un servidor default de gin utiliza un Logger y un Recuperador.
	// El logger hace logs de los requests, mientras que el recuperador permite que
	// el servidor se recupere después de un 500.
	server := gin.Default()
	// Context manages requests and JSONs
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"product": models.Product{
				ID:          1,
				Name:        "Super camera",
				Brand:       "Super brand",
				Description: "This super camera is awesome",
			},
		})
	})
	server.Run(":8080")
}
