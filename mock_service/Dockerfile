FROM golang:1.14 as builder

WORKDIR /kbc-mock-service
COPY . /kbc-mock-service

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go build -mod=vendor -v -o mock-service 
# RUN go mod tidy && go mod download
# RUN go build -mod=readonly -v -o mock-service 


FROM alpine:3

RUN apk add --no-cache ca-certificates

RUN mkdir -p /kbc-mock-service
WORKDIR /kbc-mock-service

COPY --from=builder /kbc-mock-service/mock-service  /kbc-mock-service/mock-service 
COPY --from=builder /kbc-mock-service/pubsub_credentials.json /kbc-mock-service/pubsub_credentials.json
COPY --from=builder /kbc-mock-service/config /kbc-mock-service/config

ENTRYPOINT ./mock-service
