FROM golang:1.21

WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build .

COPY . .
RUN go install github.com/cosmtrek/air@latest

EXPOSE 3490
CMD ["air"]