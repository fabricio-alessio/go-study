routes lib
https://github.com/gorilla/mux

pgAdmin
http://localhost:54321/

  PGADMIN_DEFAULT_EMAIL: "fabricio.o.alessio@gmail.com"
  PGADMIN_DEFAULT_PASSWORD: "123456"

  Get ip
  ➜  go-rest-api git:(main) ✗ docker-compose exec postgres sh
  # hostname -i
  172.24.0.2

ORM
https://gorm.io/index.html

  go get -u gorm.io/gorm
  go get gorm.io/driver/postgres
  
handlers para CORS
https://github.com/gorilla/handlers

  go get github.com/gorilla/handlers