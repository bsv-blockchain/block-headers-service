FROM golang:1.26.2 AS build-stage

ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /block-headers-service ./cmd/

FROM debian:sid-slim@sha256:5455a757130f8c31e13385c23ea357d449f7796f61a51b3695659fb5fa784b1b

WORKDIR /service

COPY --from=build-stage /block-headers-service /service/block-headers-service
COPY --from=build-stage /go/data/blockheaders.csv.gz /service/data/blockheaders.csv.gz
COPY --from=build-stage /go/data/blockheaders-testnet.csv.gz /service/data/blockheaders-testnet.csv.gz
COPY --from=build-stage /go/database/migrations /service/database/migrations

CMD ["/service/block-headers-service"]
