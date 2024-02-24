FROM golang:1.22.0-alpine3.19

RUN apk add git curl ffmpeg

WORKDIR /go/src/github.com/thootau/neko-mimu
COPY . .

RUN go get -u github.com/u2takey/ffmpeg-go
