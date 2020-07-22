FROM golang:1.14 as builder

WORKDIR /go/src/github.com/twoflower3/interview-service
COPY . .

RUN VERSION=$(git rev-parse --short HEAD) CGO_ENABLED=0 go build -ldflags "-X main.version=${VERSION}" -o /bin/interview-service cmd/*.go

FROM alpine:3.11
COPY --from=builder /bin/interview-service /bin/interview-service


ENV DEBUG="" \
    TEXTLOG="" \
    TRACE="" \
    HOST="" \
    PORT="" \
    SMTP_HOSTNAME="" \
    LOGIN="" \
    PASSWORD="" \
    SEND_MAIL="" 


CMD ["interview-service"]
