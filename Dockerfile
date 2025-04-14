FROM golang:1.23.5

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]


