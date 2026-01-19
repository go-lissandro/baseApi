FROM golang:1.23-alpine

WORKDIR /app

COPY . .

EXPOSE 3000

CMD ["./goRest"]