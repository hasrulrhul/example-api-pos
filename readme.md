# Clone Project
git@github.com:hasrulrhul/api-pos.git

# inisalisasi project
go mod init api-pos

# copy .env.example and rename .env
setting your config database

# publish vendor
go mod vendor

# run project
go run main.go
