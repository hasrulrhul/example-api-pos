# menginisalisasi projek baru
go mod init base-project-go
buat file main.go

# jalankan project
go run main.go
go build
atau

# buat file makefile kemudian
run makefile dengan perintah make dev

# install echo
go get -u github.com/labstack/echo/v4

# install gin
go get -u github.com/gin-gonic/gin

# publish vendor
go mod vendor

# package godotenv
go get package github.com/joho/godotenv

# buat file .env

# migration

# gorm 
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

# mysql
go get package github.com/go-sql-driver/mysql

go run main.go -mysql.dsn "root:@tcp(localhost)/golang"

migrate -database mysql -path /Users/hasrul/Documents/project-Go/base-project-go/db/migrations up


install library
$ brew install golang-migrate


untuk lihat versinya 
$ migrate -version

untuk melihta perintah command migrate
$ migrate -help

buat filemigration tabel users
migrate create -ext sql -dir db/migrations -seq create_users_table

Untuk update dependencies pada modul kita.
$ go mod tidy