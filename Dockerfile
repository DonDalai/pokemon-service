# Usa la imagen base de Golang 1.23.1
FROM golang:1.23.1

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum al contenedor
COPY go.mod go.sum ./

# Descarga las dependencias del proyecto
RUN go mod download

# Copia todo el código fuente del proyecto al contenedor
COPY . .

# Compila la aplicación
RUN go build -o pokemon-service

# Expone el puerto en el que se ejecutará el servicio
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./pokemon-service"]
