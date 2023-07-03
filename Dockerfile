# Dockerfile de producción

# Utiliza una imagen base de Go para producción
FROM golang:1.17-alpine

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia los archivos de la aplicación al contenedor
COPY ./lib .

# Compila la aplicación para producción
RUN go build -o main .

# Expone el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Inicia la aplicación cuando el contenedor se ejecute
CMD ["./main"]


# docker build -t mi-app-prod .
# docker run -p 8080:8080 mi-app-prod
