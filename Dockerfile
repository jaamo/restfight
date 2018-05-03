FROM golang:1.8

COPY . /go/src/app

WORKDIR /go/src/app/main

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["main"]