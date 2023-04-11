package api

import (
	"github.com/gorilla/mux"
)

// RegisterSwagger configura o Swagger na rota /swagger/
func RegisterSwagger(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), // URL para a especificação OpenAPI gerada pelo Swagger
	))
}
