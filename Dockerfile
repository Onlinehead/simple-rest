FROM golang:1.12-stretch

ADD . /go/src/github.com/onlinehead/simple-rest
WORKDIR /go/src/github.com/onlinehead/simple-rest

ARG Version

RUN GO111MODULE=on go build -o simple-rest -ldflags "-X github.com/onlinehead/simple-rest.BuildTime=`date +%Y-%m-%d:%H:%M:%S` -X github.com/onlinehead/simple-rest.AppVer=${Version}"

FROM alpine
WORKDIR /app
ENV GIN_MODE=release

RUN apk add --no-cache ca-certificates && \
  mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=0 /go/src/github.com/onlinehead/simple-rest/simple-rest /app
COPY --from=0 /go/src/github.com/onlinehead/simple-rest/postgres_migrations /app/postgres_migrations/

CMD ["/app/simple-rest"]