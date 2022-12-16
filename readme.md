Requirements:

Gorilla mux
go get -u github.com/gorilla/mux

Postgres
go get gorm.io/driver/postgres

GORM
go get -u gorm.io/gorm

To run the project:
docker-compose up

docker ps -a
docker inspect acf | grep IPAddress -- 172.19.0.2