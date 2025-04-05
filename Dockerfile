FROM golang:1.23.0

RUN mkdir /testTask

WORKDIR /testTask

COPY ./ ./

RUN go env -w GO111MODULE=on

RUN go mod download

RUN go build ./cmd/main.go

CMD ["./main"]