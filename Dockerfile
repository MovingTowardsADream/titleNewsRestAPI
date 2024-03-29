FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o titlenewsrestapi ./cmd/main.go

CMD ["./titlenewsrestapi"]