FROM golang:alpine
WORKDIR /go/src/app
COPY . .
RUN go build -o urlshortner ./cmd/api/main.go

CMD ["./urlshortner"]