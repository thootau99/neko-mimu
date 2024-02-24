FROM golang:1.22.0-alpine3.19

RUN apk add git curl ffmpeg

WORKDIR /font

RUN wget https://github.com/notofonts/noto-cjk/raw/main/Sans/OTF/Japanese/NotoSansCJKjp-Regular.otf

WORKDIR /go/src/github.com/thootau/neko-mimu
COPY . .

RUN go get -u github.com/u2takey/ffmpeg-go
