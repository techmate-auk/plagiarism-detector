################ Stage 1: Build ################

FROM golang:1.22 AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify

RUN go build -o /go/bin/app .

################ Stage 2: Runtime  ################
FROM gcr.io/distroless/static-debian12

COPY --from=builder /go/bin/app /

CMD ["/app"]
