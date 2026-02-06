# Estación Meteorológica Server

Servidor para el proyecto de Estación Meteorológica del Sirolli

## Requerimientos

- [Go >= 1.25](https://go.dev/dl/)

### Opcionales

- [air](https://github.com/air-verse/air)

## Ejecutar

Teniendo instalado Go, ejecutar:

    go build -o app .
    ./app

De forma alternativa, usando `Podman` (o en su defecto Docker):

    podman-compose up -d

## Agradecimientos

A nuestro profesor Eduardo Gomez por apoyarnos en todo el proceso de creación de la estación meteorológica.

## Referencias

* [Go net/http](https://pkg.go.dev/net/http)
* [Go Documentation](https://go.dev/doc/)
* [HTTP -Hypertext Transfer Protocol](https://www.w3.org/Protocols/Overview)

