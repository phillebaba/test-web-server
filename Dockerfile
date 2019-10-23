FROM golang:1.13-alpine AS build

COPY . /go/src/github.com/phillebaba/test-web-server/
WORKDIR /go/src/github.com/phillebaba/test-web-server/
RUN CGO_ENABLED=0 go build -o /bin/test-web-server main.go

FROM scratch
COPY --from=build /bin/test-web-server /bin/test-web-server

ENTRYPOINT ["/bin/test-web-server"]
