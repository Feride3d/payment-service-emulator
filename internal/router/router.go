package router

import (
	"net/http"

	"github.com/Feride3d/payment-service-emulator/internal/storage"
)

type Router struct { // тип роутер
	rootHandler rootHandler
}

func New(store storage.Storage) *Router { // метод new, который возвращает роутер
	return &Router{
		rootHandler: newRootHandler(store), // новый рут хэндлер, который передает стор (dependency injunction)
	}
}

func (r *Router) RootHandler() http.Handler { // метод RootHandler который возвращает http.Handler
	return r.rootHandler // должен реализовать интерфейс http.Handler
}
