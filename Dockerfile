FROM golang:1.14-alpine as builder

ENV APP_DIR=/go/src/github.com/bandoshintaro/gbot

RUN apk update
RUN apk upgrade
RUN apk --no-cache add git musl-dev

COPY . $APP_DIR
WORKDIR $APP_DIR
RUN go build -o gbot main.go gbot.go

FROM alpine:latest
COPY --from=builder /go/src/github.com/bandoshintaro/gbot/gbot /gbot
CMD ["/gbot"]
