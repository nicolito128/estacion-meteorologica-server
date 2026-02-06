# Imagen base
FROM golang:1.25-alpine3.22 AS base

ENV DIR /server
WORKDIR $DIR

FROM base AS builder

COPY go.mod go.sum *.go ./

RUN go mod download

RUN mkdir bin

RUN go build -o app .

FROM scratch AS production

WORKDIR /server

COPY --from=builder /server/app .

EXPOSE 5000

CMD ["./app", "-addr", "5000"]
