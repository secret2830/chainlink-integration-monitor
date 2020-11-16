FROM golang:1.15-alpine3.12 as builder

WORKDIR /chainlink-integration-monitor
COPY . .

RUN apk add make && make install

FROM alpine:3.12

COPY --from=builder /go/bin/chainlink-integration-monitor /usr/local/bin/

ENTRYPOINT ["chainlink-integration-monitor"]
