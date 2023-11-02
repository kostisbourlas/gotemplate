FROM golang:1.21-alpine
 
WORKDIR /app
 
COPY . /app
RUN chmod +x ./entrypoint.sh
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

RUN ["go", "install", "-mod=mod", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "-tags", "'postgres'", "github.com/golang-migrate/migrate/v4/cmd/migrate@latest"]

ENTRYPOINT ["/app/entrypoint.sh"]
