FROM golang:1.17-alpine3.14
WORKDIR /A-Golang-app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN go build -o mainfile
EXPOSE 8000

CMD ["mainfile"]