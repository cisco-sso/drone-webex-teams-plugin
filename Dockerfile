FROM golang:1.8-stretch as builder

WORKDIR /go/src/drone-spark
COPY . .

RUN go get -d github.com/drone/drone-template-lib; exit 0 
# TODO: temporary fix by adding exit 0 as it is failing in spite of -d
RUN go get github.com/pkg/errors
RUN go get github.com/aymerick/raymond
RUN go get github.com/Sirupsen/logrus
RUN go get github.com/joho/godotenv
RUN go get github.com/urfave/cli
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:3.4

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/drone-spark/drone-spark /bin/
ENTRYPOINT ["/bin/drone-spark"]
