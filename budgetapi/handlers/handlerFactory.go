package handlers

import (
	"budgetstoragelib/interfaces"
	"net/http"
)

type HandlerFactory struct {
	storage interfaces.IServer
}

func NewHandlerFactory(storage interfaces.IServer) *HandlerFactory {
	return &HandlerFactory{storage: storage}
}

func (f HandlerFactory) Load() {
	http.HandleFunc("/categories", ReturnCategoryHandlers(f.storage))
}
