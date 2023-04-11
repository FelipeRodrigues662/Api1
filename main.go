package main

import (
	api "new.com/events/Api"
	database "new.com/events/Database"
)

func main() {
	// Inicialização da conexão com o banco de dados
	database.InitDatabase()

	// Inicialização do roteador da API
	api.StartRouter()

}
