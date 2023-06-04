FROM golang

WORKDIR /app

COPY go.* /.
RUN go mod download

COPY . .
RUN go build -o /main ./src/main.go 

CMD /main