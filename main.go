//go:generate go run github.com/swaggo/swag/cmd/swag@v1.16.3 init -g main.go -d . --parseDependency --parseInternal -o docs

package main

import (
	"log"
	"net/http"

	"agent/config"
	"agent/handler"
	"agent/minio"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "agent/docs"
)

// @title		Агентный уровень
// @version		1.0
// @description	Приём запроса на конкретную страницу документа, загрузка изображения из MinIO, сегментация и отправка сегментов на транспортный уровень.
// @host		localhost:8080
// @BasePath	/
// @tag.name	agent
func main() {
	if err := minio.Init(); err != nil {
		log.Fatal("Ошибка инициализации MinIO:", err)
	}

	http.HandleFunc("/process", handler.Process)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Агентный сервис запущен на", config.ServerPort)
	log.Println("Swagger: http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(config.ServerPort, nil))
}
