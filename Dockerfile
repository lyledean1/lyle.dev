FROM golang:1.12.9 as build

COPY . /go/src/github.com/lyle.dev/

WORKDIR /go/src/github.com/lyle.dev/

RUN go get .

RUN CGO_ENABLED=0 go build -o /go/bin/lyle-dev .

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=build /go/bin/lyle-dev .

COPY ./public public

CMD ["./lyle-dev"]