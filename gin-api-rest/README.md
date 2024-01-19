// add module
go mod init github.com/falessio/gin-api-rest

// gin library
https://github.com/gin-gonic/gin

// add gin dependencies
go get -u github.com/gin-gonic/gin

// add new module to work
add use ./gin-api-rest in ../go.work

// usando pgadmin para criar base de dados
http://localhost:54321/
-> usuario e senha do docker-compose.yml
-> cria um novo servidor

// investigando IP para criar base de dados com pgadmin
➜  gin-api-rest git:(main) ✗ docker-compose exec postgres sh
# hostname -i
172.25.0.2

// gorm
go get -u gorm.io/gorm

