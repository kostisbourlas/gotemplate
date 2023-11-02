FROM golang:1.21-alpine
 
WORKDIR /app

COPY go.mod go.sum /app

RUN go mod download

RUN ["go", "install", "-mod=mod", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "-tags", "'postgres'", "github.com/golang-migrate/migrate/v4/cmd/migrate@latest"]

COPY entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

COPY . /app

ENTRYPOINT ["/app/entrypoint.sh"]
