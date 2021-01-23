FROM golang:1.15.2

WORKDIR /go/src/CloudWriteAPI

COPY . .

RUN go get ./...

CMD ["CloudWriteAPI"]