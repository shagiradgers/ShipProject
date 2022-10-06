package server

import (
	"context"
	"net/http"
	"time"
)

// Server Структура сервера
type Server struct {
	httpServer *http.Server
}

// Run запускает сервер и возвращает ошибку или nil
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,       // порт
		Handler:           handler,          // обработчик запросов
		MaxHeaderBytes:    1 << 20,          // максимальный объем хедера(1 МБ),
		ReadHeaderTimeout: 10 * time.Second, // ограничение по времени для чтения
		WriteTimeout:      10 * time.Second, // ограничение по времени для записи
	}

	return s.httpServer.ListenAndServe()
}

// ShutDown останавливает работу сервера и возвращает ошибку или nil
func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
