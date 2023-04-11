package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) {
	router := mux.NewRouter()
	RegisterHandlers(router, db)
	RegisterSwagger(router)
	http.ListenAndServe(":8000", router)
}
