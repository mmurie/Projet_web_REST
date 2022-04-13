@echo off
@echo on
go clean -cache ./...
go build -o . -v ./...
~/go/bin/swagger generate spec -o ./swagger/swagger.json
go run cmd/restserver.go
@echo off
if %ERRORLEVEL% GEQ 1 !!!!! ERROR !!!!!