FROM golang:1.22.1-alpine

WORKDIR /usr/src/app

COPY .env.docker .env

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o bin/go-fiber-boilerplate main.go

EXPOSE 3000

CMD ["bin/go-fiber-boilerplate"]