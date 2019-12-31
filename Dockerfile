FROM golang:1.13-alpine as build

WORKDIR /go/app

COPY go.mod go.mod
COPY go.sum go.sum

RUN apk add --no-cache git &&\
  go mod download && \
  go get gopkg.in/urfave/cli.v2@master && \
  go get github.com/oxequa/realize

COPY . .
RUN go build -o app

FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app

CMD ["./app"]
