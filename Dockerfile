FROM golang:1.16

WORKDIR home

COPY . .

RUN go build app.go

EXPOSE 3000

CMD ["./app"]
