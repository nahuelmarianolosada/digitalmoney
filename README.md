# digitalmoney

[EJ]docker run -d -p 33060:3306 --name digitalmoney-db -e MYSQL_ROOT_PASSWORD=mysql mysql



REQUERIMIENTOS

# Go version >= 1.16

Una vez descargado el source, instalar las dependencias con el comando:
go mod tidy

Para generar el ejecutable:
go build ./...

Para correr el proyecto:
go run main.go

Las variables de entorno se encuentran el el .env
