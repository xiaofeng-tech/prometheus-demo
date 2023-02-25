FROM golang:1.19.6-alpine3.17 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o prometheus-demo /app/cmd


FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/prometheus-demo /app/prometheus-demo

EXPOSE 8080

CMD [ "./prometheus-demo" ]
