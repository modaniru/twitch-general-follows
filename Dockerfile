FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /main ./src/main.go 

CMD ["/main"]