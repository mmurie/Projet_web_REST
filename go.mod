module Projet_web_REST

go 1.15

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/gorilla/mux v1.8.0
)

require internal/entities v1.0.0

replace internal/entities => ./internal/entities

require internal/handlers v1.0.0

replace internal/handlers => ./internal/handlers

require (
	github.com/go-openapi/analysis v0.21.3 // indirect
	github.com/go-openapi/runtime v0.23.3 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-swagger/go-swagger v0.28.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cobra v1.3.0 // indirect
	github.com/spf13/viper v1.10.1 // indirect
	go.mongodb.org/mongo-driver v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	internal/persistence v1.0.0
)

replace internal/persistence => ./internal/persistence
