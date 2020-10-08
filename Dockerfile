FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/github.com/flipper-zero/sub-sync
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/app .


FROM alpine

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/app /go/bin/app

USER appuser
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]