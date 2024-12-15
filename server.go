package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// Структура для сервера
type Server struct {
	httpServer *http.Server
}

// Метод для запуска сервера
func (s *Server) Run(port string) error {
	// Создание файла для логирования
	logFile, err := os.Create("gin.log")
	if err != nil {
		return fmt.Errorf("ошибка создания лог-файла: %v", err)
	}
	defer logFile.Close()

	// Настройка Gin для логирования в файл
	gin.DefaultWriter = logFile

	// Создание нового роута Gin
	r := gin.New()

	// Использование стандартного логгера Gin
	r.Use(gin.Logger())

	// Определение маршрутов
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Настройка HTTP-сервера
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Запуск сервера
	fmt.Printf("Запуск сервера на порту %s...\n", port)
	return s.httpServer.ListenAndServe()
}
