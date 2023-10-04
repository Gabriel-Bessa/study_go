package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"study_go/models/entity"
	"study_go/routes"
)

func main() {
	loadDotEnv()
	r := gin.Default()
	entity.ConnectToDatabase()
	routes.LoadRoutes(r)
	startApp(r)
}

func startApp(r *gin.Engine) {
	gin.SetMode(os.Getenv("GIN_MODE"))
	log.Print("Application start att port: ", os.Getenv("PORT"))
	err := r.Run("127.0.0.1:" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Erro ao iniciar a aplicação: %v", err)
	}
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}
