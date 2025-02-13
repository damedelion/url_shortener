FROM golang:1.24.0-alpine3.21

WORKDIR /url_shortener

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o url_shortener ./cmd

EXPOSE 8080

CMD [ "./url_shortener" ]
