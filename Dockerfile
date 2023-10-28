FROM golang:1.21-alpine
 
WORKDIR /app
 
COPY . /app
 
RUN go mod download
 
RUN go build -o gotemplate

CMD [ "./gotemplate" ]
