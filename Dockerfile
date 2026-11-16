# Etapa de build: compila o binário estático
FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /goRest .

# Etapa final: imagem enxuta só com o binário
FROM alpine:3.20

WORKDIR /app
COPY --from=build /goRest /app/goRest

EXPOSE 3000

CMD ["/app/goRest"]
