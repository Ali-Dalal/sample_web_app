FROM golang:1.16 AS builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR home

COPY --from=builder /build/app .

EXPOSE 3000

CMD ["./app"]
