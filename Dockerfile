FROM golang:1.12.6

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80

CMD ["app"]
