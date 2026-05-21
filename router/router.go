package router

import (
	"log"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Inicializa o router utilizando as configurações Default do gin
	r := gin.Default()

	initializeRoutes(r)

	//Inicia o server na porta 8080 por padrão
	if err := r.Run(); err != nil {
		log.Fatal("Falha em iniciar o server.")
	}
}
