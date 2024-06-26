package router

import "github.com/gin-gonic/gin"

func Inicialize () {

	// inicializa o Router utilizando as configurações Default do
	router := gin.Default()

	inicializeRoutes(router)

  router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}