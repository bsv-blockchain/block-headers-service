FROM golang@sha256:3680233e3204827fbdc66088528ae6d4b3d034f51d03a99d454f6de034888244 AS build-stage

ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /block-headers-service ./cmd/

FROM debian@sha256:9cc080028c43b27d2074d63a5f9caf7166d731494965616c1a6d2827a004585c

WORKDIR /service

COPY --from=build-stage /block-headers-service /service/block-headers-service
COPY --from=build-stage /go/data/blockheaders.csv.gz /service/data/blockheaders.csv.gz
COPY --from=build-stage /go/data/blockheaders-testnet.csv.gz /service/data/blockheaders-testnet.csv.gz
COPY --from=build-stage /go/database/migrations /service/database/migrations

CMD ["/service/block-headers-service"]
