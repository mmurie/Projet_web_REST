module Projet_web_REST

go 1.15

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
)

require internal/entities v1.0.0

replace internal/entities => ./internal/entities

require internal/handlers v1.0.0

replace internal/handlers => ./internal/handlers

require internal/persistence v1.0.0

replace internal/persistence => ./internal/persistence
