package main

import "github.com/gin-gonic/gin"

func main() {
	// print("A galera da live é braba e vai conseguir trampo de 200k/mês!")

	// inicializa o Router utilizando as configurações Default do
	router := gin.Default()

	// definindo uma nova Rota
	// "/ping" => realtive Path, func(c *gin.Context) => handlers
	router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}