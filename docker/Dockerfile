FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags "-X main.version=1.0.0" -o golang-cli-template

CMD [ "./golang-cli-template" ]
