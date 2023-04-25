FROM golang:1.18.8-alpine3.16 AS builder

RUN mkdir -p /go/src/jpagent
COPY ./ /go/src/jpagent
WORKDIR /go/src/jpagent
ENV  GO111MODULE=on
RUN cd /go/src/jpagent && go build -o JPAgent -mod vendor

EXPOSE 80

FROM alpine:3.11.6
COPY --from=builder /go/src/jpagent/JPAgent /go/src/jpagent/config.yaml ./
ENTRYPOINT ["./JPAgent"]
