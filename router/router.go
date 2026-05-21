package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Inicializa o router utilizando as configurações Default do gin
	r := gin.Default()

	//Define um simples GET de endpoint
	r.GET("/ping", func(c *gin.Context) {
		//retorna uma resposta json
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	//Inicia o server na porta 8080 por padrão
	if err := r.Run(); err != nil {
		log.Fatal("Falha em iniciar o server.")
	}
}
