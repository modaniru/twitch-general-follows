FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./src ./src

RUN go build -o /main ./src/main.go 

CMD ["/main"]