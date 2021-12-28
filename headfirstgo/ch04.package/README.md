mkdir greeting
cd greeting
go mod init greeting
vi greeting.go
cd ..
mkdir testapp
cd testapp
go mod testapp
vi main.go
go mod edit -replace greeting=../greeting
go mod edit -replace keyboard=../keyboard
go mod tidy
go run testapp
