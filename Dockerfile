FROM golang:1.21-alpine
 
WORKDIR /app
 
COPY . /app

RUN ["go", "install", "-mod=mod", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o gotemplate" -command="./gotemplate"
