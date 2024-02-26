FROM golang:alpine AS builder

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o /l0 cmd/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow

WORKDIR /app
COPY --from=builder /l0 /app/l0

ADD /templates /app/templates

CMD ["./l0"]